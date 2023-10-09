package doublyLinkedList

import (
	"fmt"
	"log"
	"strings"
)

type Node[T any] struct {
	next *Node[T]
	prev *Node[T]
	val  T
}

type DoublyLinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func NewNode[T any](data T, prev *Node[T], next *Node[T]) *Node[T] {
	return &Node[T]{
		next: next,
		prev: prev,
		val:  data,
	}
}

func (dl *DoublyLinkedList[T]) Size() int {
	return dl.size
}

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

func (dl *DoublyLinkedList[T]) IsEmpty() bool {
	return dl.size == 0
}

func (dl *DoublyLinkedList[T]) Add(elem T) {
	dl.AddLast(elem)
}

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
