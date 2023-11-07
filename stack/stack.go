// Package stack implements the stack data structure.
// uses my doublyLinkedList package for the stack
package stack

import (
	"log"

	"github.com/smekuria1/GoSlow/doublyLinkedList"
)

// Stack represents the stack data structure.
// implemented with a doubly linked list.
type Stack[T any] struct {
	list *doublyLinkedList.DoublyLinkedList[T]
}

// NewStack creates a new instance of Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		list: doublyLinkedList.NewDoublyLinkedList[T](),
	}
}

// Push pushes an element to the stack.
func (s *Stack[T]) Push(val T) {
	s.list.Add(val)
}

// Pop pops an element from the stack.
func (s *Stack[T]) Pop() T {
	if s.list.IsEmpty() {
		log.Fatal("Stack is empty")
	}
	return s.list.RemoveAt(s.list.Size() - 1)
}

// Peek returns the top element of the stack.
func (s *Stack[T]) Peek() T {
	if s.list.IsEmpty() {
		log.Fatal("Stack is empty")
	}
	return s.list.PeekLast()
}

// Size returns the size of the stack.
func (s *Stack[T]) Size() int {
	return s.list.Size()
}

// IsEmpty returns true if the stack is empty, false otherwise.
func (s *Stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

// Clear clears the stack.
func (s *Stack[T]) Clear() {
	s.list.Clear()
}

// ToString returns a string representation of the stack.
func (s *Stack[T]) ToString() string {
	s.list.Reverse()
	return s.list.ToString()
}

// Contains returns true if the stack contains the element, false otherwise.
func (s *Stack[T]) Contains(val T) bool {
	return s.list.Contains(val)
}
