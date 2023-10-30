// Package hashtable implements a hash table data structure using separate chaining for collision resolution.
// uses my doublyLinkedList package for the buckets
package hashtable

import (
	"fmt"
	"strings"

	"github.com/smekuria1/GoSlow/doublyLinkedList"
)

// Entry represents a key-value pair.
type Entry[K, V comparable] struct {
	key   K
	value V
	Hash  int
}

// NewEntry creates a new key-value pair entry by hashing the key that is of type K comparable
func NewEntry[K, V comparable](key K, value V) *Entry[K, V] {
	h := Hash(key)
	return &Entry[K, V]{key: key, value: value, Hash: h}
}

// Equals checks if two entries are equal
func (e *Entry[K, V]) Equals(other *Entry[K, V]) bool {
	if e.Hash != other.Hash {
		return false
	}
	return e.key == other.key && e.value == other.value
}

// ToString returns a string representation of the entry
func (e Entry[K, V]) ToString() string {
	return fmt.Sprintf("%v:%v", e.key, e.value)
}

// HashTable represents a hash table data structure
type HashTable[K, V comparable] struct {
	maxLoadFactor float64
	capacity      int
	size          int
	threshold     int
	table         []doublyLinkedList.DoublyLinkedList[Entry[K, V]]
}

// NewHashTable creates a new hash table with the given capacity(optional) and load factor(optional)
// if no arguments are given, the default values of 0.75 and 16 are used
func NewHashTable[K, V comparable](capacity int, maxLoadFactor float64) *HashTable[K, V] {
	if capacity <= 0 {
		capacity = 16
	}
	if maxLoadFactor <= 0 {
		maxLoadFactor = 0.75
	}
	return &HashTable[K, V]{
		maxLoadFactor: maxLoadFactor,
		capacity:      capacity,
		size:          0,
		threshold:     int(float64(capacity) * maxLoadFactor),
		table:         make([]doublyLinkedList.DoublyLinkedList[Entry[K, V]], capacity),
	}
}

// Size returns the number of elements in the hash table
func (h *HashTable[K, V]) Size() int {
	return h.size
}

// IsEmpty returns true if the hash table is empty
func (h *HashTable[K, V]) IsEmpty() bool {
	return h.size == 0
}

// Clear removes all elements from the hash table
func (h *HashTable[K, V]) Clear() {
	h.table = make([]doublyLinkedList.DoublyLinkedList[Entry[K, V]], h.capacity)
	h.size = 0
}

// normalizeIndex converts a hash code into a valid bucket index
func (h *HashTable[K, V]) normalizeIndex(keyHash int) int {
	return keyHash % h.capacity
}

// ContainsKey checks if the hash table contains the given key
func (h *HashTable[K, V]) ContainsKey(key K) bool {
	return h.Has(key)
}

// Has checks if the hash table contains the given key
func (h *HashTable[K, V]) Has(key K) bool {
	bucketIndex := h.normalizeIndex(Hash(key))
	return h.bucketSeekEntry(bucketIndex, key) != nil
}

// Put inserts the given key-value pair into the hash table
func (h *HashTable[K, V]) Put(key K, value V) V {
	return h.Insert(key, value)
}

// Add inserts the given key-value pair into the hash table
func (h *HashTable[K, V]) Add(key K, value V) V {
	return h.Insert(key, value)
}

// Insert inserts the given key-value pair into the hash table
func (h *HashTable[K, V]) Insert(key K, value V) V {
	entry := NewEntry[K, V](key, value)
	// fmt.Println(entry.ToString())
	bucketIndex := h.normalizeIndex(entry.Hash)
	// fmt.Printf("bucketIndex: %v\n", bucketIndex)
	return *h.bucketInsertEntry(bucketIndex, entry)
}

// Get returns the value associated with the given key return nil if the key is not found
func (h *HashTable[K, V]) Get(key K) *V {
	bucketIndex := h.normalizeIndex(Hash(key))
	entry := h.bucketSeekEntry(bucketIndex, key)
	if entry != nil {
		return &entry.value
	}
	return nil
}

// bucketSeekEntry searches for an entry with the given key in the given bucket
func (h *HashTable[K, V]) bucketSeekEntry(bucketIndex int, key K) *Entry[K, V] {

	bucket := h.table[bucketIndex]
	for i := 0; i < bucket.Size(); i++ {
		entry := bucket.Get(i)
		if entry.key == key {
			return &entry
		}
	}
	// if we reach here, the entry was not found empty bucket
	return nil

}

// bucketInsertEntry inserts the given entry in the given bucket only if the entry does not already exist in the bucket
// if the entry already exists, the value is updated and the old value is returned
func (h *HashTable[K, V]) bucketInsertEntry(bucketIndex int, entry *Entry[K, V]) *V {
	bucket := h.table[bucketIndex]

	existingEntry := h.bucketSeekEntry(bucketIndex, entry.key)
	if existingEntry == nil {
		if h.size+1 > h.threshold {
			h.resizeTable()
		}
		bucket.Add(*entry)
		h.size++
		h.table[bucketIndex] = bucket
	} else {
		oldValue := existingEntry.value
		existingEntry.value = entry.value
		return &oldValue
	}
	return &entry.value
}

// resizeTable resizes the hash table to double its current capacity
func (h *HashTable[K, V]) resizeTable() {
	h.capacity *= 2
	h.threshold = int(float64(h.capacity) * h.maxLoadFactor)

	newTable := make([]doublyLinkedList.DoublyLinkedList[Entry[K, V]], h.capacity)

	// rehash all existing entries
	for i := 0; i < len(h.table); i++ {
		if h.table[i].Size() > 0 {
			for j := 0; j < h.table[i].Size(); j++ {
				entry := h.table[i].Get(j)
				bucketIndex := h.normalizeIndex(entry.Hash)
				newTable[bucketIndex].Add(entry)
			}

		}
		h.table[i].Clear()
	}
	h.table = newTable

}

// Remove removes the entry with the given key from the hash table
func (h *HashTable[K, V]) Remove(key K) *V {
	bucketIndex := h.normalizeIndex(Hash(key))
	return h.bucketRemoveEntry(bucketIndex, key)
}

// bucketRemoveEntry removes the entry with the given key from the given bucket
func (h *HashTable[K, V]) bucketRemoveEntry(bucketIndex int, key K) *V {
	entry := h.bucketSeekEntry(bucketIndex, key)
	if entry != nil {
		bucket := h.table[bucketIndex]
		bucket.RemoveVal(*entry)
		h.size--
		h.table[bucketIndex] = bucket
		return &entry.value
	}
	return nil
}

// Keys returns a slice of all the keys in the hash table
func (h *HashTable[K, V]) Keys() []K {
	var keys []K
	for i := 0; i < h.capacity; i++ {
		if h.table[i].Size() > 0 {
			for j := 0; j < h.table[i].Size(); j++ {
				keys = append(keys, h.table[i].Get(j).key)
			}
		}
	}
	return keys
}

// Values returns a slice of all the values in the hash table
func (h *HashTable[K, V]) Values() []V {
	var values []V
	for i := 0; i < h.capacity; i++ {
		if h.table[i].Size() > 0 {
			for j := 0; j < h.table[i].Size(); j++ {
				values = append(values, h.table[i].Get(j).value)
			}
		}
	}
	return values
}

// ToString returns a string representation of the hash table
func (h *HashTable[K, V]) ToString() string {
	if h.IsEmpty() {
		return "{}"
	}

	var builder strings.Builder
	builder.WriteString("{")
	for i := 0; i < h.capacity; i++ {
		if h.table[i].Size() > 0 {
			for j := 0; j < h.table[i].Size(); j++ {
				builder.WriteString(h.table[i].Get(j).ToString())
				if i != h.capacity-1 || j != h.table[i].Size()-1 {
					builder.WriteString(", ")
				}
			}
		}
	}
	builder.WriteString("}")
	return builder.String()
}
