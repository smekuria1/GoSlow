// Package doublyLinkedList provides an implementation of a doubly linked list
// data structure.
package doublyLinkedList

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// Node represents a node in the doubly linked list.
type Node[T any] struct {
	next *Node[T]
	prev *Node[T]
	val  T
}

// DoublyLinkedList represents the doubly linked list data structure.
type DoublyLinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

// NewDoublyLinkedList creates a new instance of DoublyLinkedList.
func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		size: 0,
		head: nil,
		tail: nil,
	}
}

// NewNode creates a new instance of Node.
func NewNode[T any](val T, prev *Node[T], next *Node[T]) *Node[T] {
	return &Node[T]{
		next: next,
		prev: prev,
		val:  val,
	}
}

// Size returns the size of the doubly linked list.
func (dl *DoublyLinkedList[T]) Size() int {
	return dl.size
}

// Clear clears the doubly linked list.
func (dl *DoublyLinkedList[T]) Clear() {
	trav := dl.head
	for trav != nil {
		next := trav.next
		trav.prev = nil
		trav.next = nil
		trav = next
	}
	dl.head = nil
	dl.tail = nil
	trav = nil
	dl.size = 0
}

// IsEmpty checks if the doubly linked list is empty.
func (dl *DoublyLinkedList[T]) IsEmpty() bool {
	return dl.size == 0
}

// Add adds an element to the doubly linked list.
func (dl *DoublyLinkedList[T]) Add(elem T) {
	dl.AddLast(elem)
}

// AddFirst adds an element to the beginning of the doubly linked list.
func (dl *DoublyLinkedList[T]) AddFirst(elem T) {
	if dl.IsEmpty() {
		dl.head = NewNode[T](elem, nil, nil)
		dl.tail = NewNode[T](elem, nil, nil)
	} else {
		dl.head.prev = NewNode[T](elem, nil, dl.head)
		dl.head = dl.head.prev
	}

	dl.size += 1

}

// AddLast adds an element to the end of the doubly linked list.
func (dl *DoublyLinkedList[T]) AddLast(elem T) {
	if dl.IsEmpty() {
		newNode := NewNode[T](elem, dl.tail, nil)

		dl.head = newNode
		dl.tail = newNode
	} else {
		newNode := NewNode[T](elem, dl.tail, nil)
		dl.tail.next = newNode
		dl.tail = dl.tail.next
	}
	dl.size += 1
}

// AddAt adds an element at a specific index in the doubly linked list.
func (dl *DoublyLinkedList[T]) AddAt(index int, val T) {
	if index < 0 || index > dl.size {
		log.Fatal("Illegal Index")
	}
	if index == 0 {
		dl.AddFirst(val)
		return
	}

	if index == dl.size {
		dl.AddLast(val)
		return
	}

	temp := dl.head

	for i := 0; i < index-1; i++ {
		temp = temp.next
	}

	newNode := NewNode[T](val, temp, temp.next)
	temp.next.prev = newNode
	temp.next = newNode

	dl.size += 1

}

// ToString returns a string representation of the doubly linked list.
func (dl *DoublyLinkedList[T]) ToString() string {
	if dl.size == 0 {
		return "[]"
	}
	var builder strings.Builder
	builder.WriteString("[")
	trav := dl.head

	for trav != nil {
		builder.WriteString(fmt.Sprintf("%v", trav.val))
		if trav.next != nil {
			builder.WriteString(",")
		}
		trav = trav.next

	}
	builder.WriteString("]")
	return builder.String()
}

// PeekFirst Check the value of the first node if exists, 0(1)
func (dl *DoublyLinkedList[T]) PeekFirst() T {
	if dl.IsEmpty() {
		log.Fatal("Empty list")
	}

	return dl.head.val
}

// PeekLast Check the value of the last node if it exists, 0(1)
func (dl *DoublyLinkedList[T]) PeekLast() T {
	if dl.IsEmpty() {
		log.Fatal("Empty list")
	}

	return dl.tail.val
}

// RemoveFirst removes the first value at head of the LinkedList
func (dl *DoublyLinkedList[T]) RemoveFirst() T {
	if dl.IsEmpty() {
		log.Fatal("Empty List")
	}

	val := dl.head.val
	dl.head = dl.head.next

	dl.size -= 1

	if dl.IsEmpty() {
		dl.tail = nil
	} else {
		dl.head.prev = nil
	}

	return val
}

// RemoveLast removes the last value at the tail of the Linked List, O(1)
func (dl *DoublyLinkedList[T]) RemoveLast() T {
	if dl.IsEmpty() {
		log.Fatal("Empty List")
	}

	val := dl.tail.val
	dl.tail = dl.tail.prev
	dl.size -= 1
	if dl.IsEmpty() {
		dl.head = nil
	} else {
		dl.tail.next = nil
	}

	return val
}

// Remove removes an aribitrary node from the Linked List, 0(1) INTERNAL
func (dl *DoublyLinkedList[T]) remove(node Node[T]) T {
	if node.prev == nil {
		return dl.RemoveFirst()
	}
	if node.next == nil {
		return dl.RemoveLast()
	}
	node.next.prev = node.prev
	node.prev.next = node.next

	val := node.val

	// Not cleaning up memory I dont fully understand the GO GC and how
	// it deals with freeing up the node value when it gets deleted
	//node.val = nil

	node.prev = nil
	node.next = nil

	dl.size -= 1
	return val
}

// RemoveAt removes a node at particular index O(n)
func (dl *DoublyLinkedList[T]) RemoveAt(index int) T {

	if index < 0 || index >= dl.size {
		log.Fatal("Illegal Index")
	}

	var i int
	var trav *Node[T]

	if index < dl.size/2 {
		trav = dl.head
		for i = 0; i != index; i++ {
			trav = trav.next
		}
	} else {
		trav = dl.head
		for i := dl.size - 1; i != index; i-- {
			trav = trav.prev
		}
	}

	return dl.remove(*trav)
}

// Since I am not enforcing types for added items i am writing a catch all compare function
// janky chat gpt code sry
// type Comparable interface {
// 	Compare(interface{}) int
// }
// func compareValues(a, b interface{}) (int, error) {
//     // Use reflection to get the types of a and b.
//     typeOfA := reflect.TypeOf(a)
//     typeOfB := reflect.TypeOf(b)
//     // Check if the types are comparable.
//     if typeOfA != typeOfB {
//         return 0, fmt.Errorf("types are not comparable: %s and %s", typeOfA, typeOfB)
//     }
//     // Check if the types implement the Comparable interface.
//     compareMethod := reflect.ValueOf(a).MethodByName("Compare")
//     if !compareMethod.IsValid() {
//         return 0, fmt.Errorf("Compare method not found for type %s", typeOfA)
//     }
//     // Call the Compare method on the values.
//     result := compareMethod.Call([]reflect.Value{reflect.ValueOf(b)})
//     // Extract the result as an int.
//     if len(result) > 0 {
//         if result[0].Kind() == reflect.Int {
//             return int(result[0].Int()), nil
//         }
//     }
//     return 0, fmt.Errorf("comparison result not as expected")
// }

// slightly better compare using reflect
func isEqual(x, y interface{}) bool {
	return reflect.DeepEqual(x, y)
}

// // RemoveVal removes a node by its value O(n)
// func (dl *DoublyLinkedList[T]) RemoveVal(val T) bool {
// 	trav := dl.head

// 	// Not sure about allowing nil types to be in the list dont fully understand generics yet
// 	if reflect.TypeOf(val) == nil {
// 		for trav = dl.head; trav != nil; trav = trav.next {
// 			if reflect.TypeOf(trav.val) == nil {
// 				dl.remove(*trav)
// 				return true
// 			}

// 		}
// 	} else {
// 		for trav = dl.head; trav != nil; trav = trav.next {
// 			if isEqual(val, trav.val) {
// 				dl.remove(*trav)
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// RemoveVal removes a node by its value O(n) relies on reflect
func (dl *DoublyLinkedList[T]) RemoveVal(val T) bool {
	trav := dl.head

	// Iterate through the list to find the node with the specified value.
	for trav != nil {
		if isEqual(val, trav.val) { // Use direct comparison to check if values are equal.
			dl.remove(*trav)
			return true
		}
		trav = trav.next
	}

	return false // Value not found in the list.
}

// IndexOf find the index of a particular value in the linked list, O(n)
func (dl *DoublyLinkedList[T]) IndexOf(elem T) int {
	index := 0
	trav := dl.head

	for trav != nil {
		if isEqual(trav.val, elem) {
			return index
		}
		index++
		trav = trav.next
	}
	return -1
}

// Contains check if the value is contained within the linked list
func (dl *DoublyLinkedList[T]) Contains(elem T) bool {
	return dl.IndexOf(elem) != -1
}

// Get returns the value at a particular index, O(n)
func (dl *DoublyLinkedList[T]) Get(index int) T {
	if index < 0 || index >= dl.size {
		log.Fatal("Illegal Index")
	}

	var i int
	var trav *Node[T]

	if index < dl.size/2 {
		trav = dl.head
		for i = 0; i != index; i++ {
			trav = trav.next
		}
	} else {
		trav = dl.head
		for i := dl.size - 1; i != index; i-- {
			trav = trav.prev
		}
	}

	return trav.val
}

// Reverse reverses the linked list in place
func (dl *DoublyLinkedList[T]) Reverse() {
	trav := dl.head
	var temp *Node[T]

	for trav != nil {
		temp = trav.prev
		trav.prev = trav.next
		trav.next = temp
		trav = trav.prev
	}

	if temp != nil {
		dl.head = temp.prev
	}
}
