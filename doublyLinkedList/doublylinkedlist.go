// Package doublyLinkedList provides an implementation of a doubly linked list
// data structure.
package doublyLinkedList

import (
	"fmt"
	"log"
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
func NewNode[T any](data T, prev *Node[T], next *Node[T]) *Node[T] {
	return &Node[T]{
		next: next,
		prev: prev,
		val:  data,
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
		dl.head = NewNode[T](elem, nil, nil)
		dl.tail = NewNode[T](elem, nil, nil)
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
		log.Printf("Trav %v", trav)
		log.Printf("Trav prev: %v", trav.prev)
		log.Printf("Trav next: %v", trav.next)

		trav = trav.next

	}
	builder.WriteString("]")
	return builder.String()
}
