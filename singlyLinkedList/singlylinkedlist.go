package singlylinkedlist

import (
	"fmt"
	"strings"
)

// Node represents a node in the singly linked list.
type Node[T any] struct {
	next *Node[T]
	val  T
}

// SinglyLinkedList represents the singly linked list data structure.
type SinglyLinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

// NewSinglyLinkedList creates a new instance of SinglyLinkedList.
func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		size: 0,
		head: nil,
		tail: nil,
	}
}

// NewNode creates a new instance of Node.
func NewNode[T any](val T, next *Node[T]) *Node[T] {
	return &Node[T]{
		next: next,
		val:  val,
	}
}

// Size returns the size of the singly linked list.
func (sl *SinglyLinkedList[T]) Size() int {
	return sl.size
}

// Clear clears the singly linked list.
func (sl *SinglyLinkedList[T]) Clear() {
	trav := sl.head
	for trav != nil {
		next := trav.next
		trav.next = nil
		trav = next
	}
	sl.head = nil
	sl.tail = nil
	trav = nil
	sl.size = 0
}

// IsEmpyty checks if the singly linked list is empty
func (sl *SinglyLinkedList[T]) IsEmpty() bool {
	return sl.size == 0
}

// Add adds an element to the end singly linked list
func (sl *SinglyLinkedList[T]) Add(elem T) {
	if sl.IsEmpty() {
		newNode := NewNode[T](elem, nil)

		sl.head = newNode
	} else {
		newNode := NewNode[T](elem, nil)
		head := sl.head
		for head.next != nil {
			head = head.next
		}
		head.next = newNode

	}
	sl.size += 1
}

// ToString returns a string representation of the singly linked list
func (sl *SinglyLinkedList[T]) ToString() string {
	if sl.IsEmpty() {
		return "[]"
	}
	var builder strings.Builder
	builder.WriteString("[")
	trav := sl.head
	for trav != nil {
		builder.WriteString(fmt.Sprintf("%v", trav.val))
		if trav.next != nil {
			builder.WriteString(", ")
		}
		trav = trav.next

	}
	builder.WriteString("]")
	return builder.String()

}

// AddFirst adds an element to the beginning of the singly linked list returns new head
func (sl *SinglyLinkedList[T]) AddFirst(elem T) Node[T] {
	if sl.IsEmpty() {
		sl.Add(elem)
	} else {
		newHead := NewNode[T](elem, sl.head)
		sl.head = newHead
	}

	sl.size += 1

	return *sl.head
}

// AddAt adds an element at a specific index in the singly linked list
func (sl *SinglyLinkedList[T]) AddAt(index int, val T) error {
	if index < 0 || index > sl.size {
		return fmt.Errorf("illegal index")
	}
	if index == 0 {
		sl.AddFirst(val)
		return nil
	}

	temp := sl.head

	for i := 0; i < index-1; i++ {
		temp = temp.next
	}

	newNode := NewNode[T](val, temp.next)
	temp.next = newNode

	sl.size += 1
	return nil

}

// Remove removes the last element from the singly linked list returns the removed element
func (sl *SinglyLinkedList[T]) Remove() (*T, error) {
	if sl.IsEmpty() {
		return nil, fmt.Errorf("empty list")
	}
	if sl.size == 1 {
		removed := sl.head.val
		sl.head = nil
		sl.tail = nil
		sl.size -= 1
		return &removed, nil
	}
	temp := sl.head
	for temp.next.next != nil {
		temp = temp.next
	}
	removed := temp.next.val
	temp.next = nil
	sl.size -= 1
	return &removed, nil
}
