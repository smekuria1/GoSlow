package hashtable

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/smekuria1/GoSlow/doublyLinkedList"
)

func TestNewEntry(t *testing.T) {
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name string
		args args
		want *Entry[int, int]
	}{
		{"TestNewEntry", args{1, 1}, &Entry[int, int]{key: 1, value: 1, Hash: Hash(1)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEntry[int, int](tt.args.key, tt.args.value); !reflect.DeepEqual(got.Hash, tt.want.Hash) {
				t.Errorf("NewEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntry_Equals(t *testing.T) {
	type args struct {
		other *Entry[int, int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"TestEntry_Equals", args{&Entry[int, int]{key: 1, value: 1, Hash: Hash(1)}}, true},
		{"TestEntry_Equals", args{&Entry[int, int]{key: 2, value: 2, Hash: Hash(2)}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEntry[int, int](1, 1)
			if got := e.Equals(tt.args.other); got != tt.want {
				t.Errorf("Entry.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHashTable(t *testing.T) {
	type args struct {
		capacity      int
		maxLoadFactor float64
	}
	type K int
	type V int
	tests := []struct {
		name string
		args args
		want *HashTable[K, V]
	}{
		{"TestNewHashTable", args{16, 0.75}, &HashTable[K, V]{maxLoadFactor: 0.75, capacity: 16, size: 0, threshold: 12, table: make([]doublyLinkedList.DoublyLinkedList[Entry[K, V]], 16)}},
		{"TestNewHashTable", args{0, 0.75}, &HashTable[K, V]{maxLoadFactor: 0.75, capacity: 16, size: 0, threshold: 12, table: make([]doublyLinkedList.DoublyLinkedList[Entry[K, V]], 16)}},
		{"TestNewHashTable", args{16, 0}, &HashTable[K, V]{maxLoadFactor: 0.75, capacity: 16, size: 0, threshold: 12, table: make([]doublyLinkedList.DoublyLinkedList[Entry[K, V]], 16)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashTable[K, V](tt.args.capacity, tt.args.maxLoadFactor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Put(t *testing.T) {
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name string
		args args
		want *HashTable[int, int]
	}{
		{
			"TestHashTable_Put",
			args{300, 300},
			get_setup_hashtable(),
		},
		{
			"TestHashTable_Put",
			args{5, 5},
			get_setup_hashtable(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashTable[int, int](16, 0.75)
			h.Put(tt.args.key, tt.args.value)
			if got := h.Get(tt.args.key); got == nil {
				t.Errorf("HashTable.Get() = %v, want %v", got, tt.args.value)
			} else {
				if *got != tt.args.value {
					t.Errorf("HashTable.Get() = %v, want %v", *got, tt.args.value)
				}
			}
		})
	}
}

func TestHashTable_Add(t *testing.T) {
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name string
		args args
		want *HashTable[int, int]
	}{
		{
			"TestHashTable_Add",
			args{300, 300},
			get_setup_hashtable(),
		},
		{
			"TestHashTable_Add",
			args{5, 5},
			get_setup_hashtable(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := get_setup_hashtable()
			h.Add(tt.args.key, tt.args.value)
			if got := h.Get(tt.args.key); got == nil {
				t.Errorf("HashTable.Get() = %v, want %v", got, tt.args.value)
			} else {
				if *got != tt.args.value {
					t.Errorf("HashTable.Get() = %v, want %v", *got, tt.args.value)
				}
			}
		})
	}
}

// generate random numbers and insert them into the hashtable
func get_setup_hashtable() *HashTable[int, int] {
	h := NewHashTable[int, int](16, 0.75)
	for i := 0; i < 11; i++ {
		h.Add(i, i)
	}
	// create a collision
	h.Add(12, 12)
	h.Add(28, 28)
	return h
}

func TestHashTable_Get(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"TestHashTable_Get", args{5}, 5},
		{"TestHashTable_Get", args{12}, 12},
		{"TestHashTable_Get", args{13}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := get_setup_hashtable()
			// check if got is nil or not
			if got := h.Get(tt.args.key); got == nil && tt.want != nil {
				t.Errorf("HashTable.Get() = %v, want %v", got, tt.want)
			} else {
				if got == nil && tt.want == nil {
					return
				}
				if *got != tt.want {
					t.Errorf("HashTable.Get() = %v, want %v", *got, tt.want)
				}
			}
		})
	}
}

func TestHashTable_Has(t *testing.T) {
	type args struct {
		key int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"TestHashTable_Has", args{5}, true},
		{"TestHashTable_Has", args{12}, true},
		{"TestHashTable_Has", args{13}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := get_setup_hashtable()
			if got := h.Has(tt.args.key); got != tt.want {
				t.Errorf("HashTable.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_ContainsKey(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"TestHashTable_ContainsKey", args{5}, true},
		{"TestHashTable_ContainsKey", args{12}, true},
		{"TestHashTable_ContainsKey", args{13}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := get_setup_hashtable()
			if got := h.ContainsKey(tt.args.key); got != tt.want {
				t.Errorf("HashTable.ContainsKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Remove(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		args args
		want *HashTable[int, int]
	}{
		{"TestHashTable_Remove", args{5}, get_setup_hashtable()},
		{"TestHashTable_Remove", args{12}, get_setup_hashtable()},
		{"TestHashTable_Remove", args{13}, get_setup_hashtable()},
		{"TestHashTable_Remove", args{28}, get_setup_hashtable()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := get_setup_hashtable()
			h.Remove(tt.args.key)
			if got := h.Get(tt.args.key); got != nil {
				for _, v := range h.table {
					fmt.Println(v.ToString())
				}
				t.Errorf("HashTable.Get() = %v, want %v", got, nil)
			}
		})
	}
}

func TestHashTable_Size(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"TestHashTable_Size", 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := get_setup_hashtable()
			if got := h.Size(); got != tt.want {
				t.Errorf("HashTable.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"TestHashTable_IsEmpty", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := get_setup_hashtable()
			if got := h.IsEmpty(); got != tt.want {
				t.Errorf("HashTable.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Clear(t *testing.T) {
	tests := []struct {
		name string
		want *HashTable[int, int]
	}{
		{"TestHashTable_Clear", get_setup_hashtable()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := get_setup_hashtable()
			h.Clear()
			if got := h.Size(); got != 0 {
				t.Errorf("HashTable.Size() = %v, want %v", got, 0)
			}
		})
	}
}

func TestHashTable_Keys(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{"TestHashTable_Keys", []int{3, 4, 10, 8, 28, 6, 12, 0, 5, 9, 2, 7, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := get_setup_hashtable()
			if got := h.Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashTable.Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Values(t *testing.T) {
	tests := []struct {
		name      string
		want      []int
		hashtable *HashTable[int, int]
	}{
		{"TestHashTable_Values", []int{3, 4, 10, 8, 28, 6, 12, 0, 5, 9, 2, 7, 1}, get_setup_hashtable()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hashtable.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashTable.Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_ToString(t *testing.T) {
	tests := []struct {
		name      string
		want      string
		hashtable *HashTable[int, int]
	}{
		{"TestHashTable_ToString", "{3:3, 4:4, 10:10, 8:8, 28:28, 6:6, 12:12, 0:0, 5:5, 9:9, 2:2, 7:7, 1:1, }", get_setup_hashtable()},
		{"TestNewHashTable", "{}", NewHashTable[int, int](16, 0.75)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hashtable.ToString(); got != tt.want {
				t.Errorf("HashTable.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
