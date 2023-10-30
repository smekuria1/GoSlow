// Package queue implements the queue data structure.
// uses my doublyLinkedList package for the queue
package queue

import (
	"log"

	"github.com/smekuria1/GoSlow/doublyLinkedList"
)

// Queue represents the queue data structure.
// implemented with a doubly linked list.
type Queue[T any] struct {
	list *doublyLinkedList.DoublyLinkedList[T]
}

// NewQueue creates a new instance of Queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		list: doublyLinkedList.NewDoublyLinkedList[T](),
	}
}

// Size returns the size of the queue.
func (q *Queue[T]) Size() int {
	return q.list.Size()
}

// IsEmpty returns true if the queue is empty, false otherwise.
func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

// Clear clears the queue.
func (q *Queue[T]) Clear() {
	q.list.Clear()
}

// Enqueue adds an element to the queue.
func (q *Queue[T]) Enqueue(val T) {
	q.list.Add(val)
}

// Dequeue removes an element from the queue.
func (q *Queue[T]) Dequeue() T {
	if q.list.IsEmpty() {
		log.Fatal("Queue is empty")
	}
	return q.list.RemoveAt(0)
}

// Peek returns the first element of the queue.
func (q *Queue[T]) Peek() T {
	if q.list.IsEmpty() {
		log.Fatal("Queue is empty")
	}
	return q.list.PeekFirst()
}

// ToString returns a string representation of the queue.
func (q *Queue[T]) ToString() string {
	return q.list.ToString()
}
