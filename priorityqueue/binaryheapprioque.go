// Package priorityqueue implements a priority queue.
// uses my dynamicArray package for the heap
package priorityqueue

import (
	"reflect"

	darray "github.com/smekuria1/GoSlow/darray"
)

// BinaryHeapPQ is a priority queue implemented using a binary heap.
// The type T must be comparable.
type BinaryHeapPQ[T comparable] struct {
	heap *darray.DynamicArray[T]
}

// NewBinaryHeapQP returns a new instance of BinaryHeapPQ[T] with the given capacity.
// If no capacity is provided, the default capacity of 10 is used.
func NewBinaryHeapPQ[T comparable](args ...int) *BinaryHeapPQ[T] {
	if len(args) != 0 {
		arr := darray.NewDynamicArray[T](args[0])
		return &BinaryHeapPQ[T]{
			heap: arr,
		}
	}

	return &BinaryHeapPQ[T]{
		heap: darray.NewDynamicArray[T](),
	}
}

// TODO: refactor this method
//
//	doesnt maintain heap invariant
//
// NewHeapifiedBinaryPQ returns a new instance of BinaryHeapPQ[T] with the given elements.
// The elements are heapified using the binary heap algorithm.
func (pq *BinaryHeapPQ[T]) NewHeapifiedBinaryPQ(elems []int) *BinaryHeapPQ[T] {
	heapSize := len(elems)
	arr := darray.NewDynamicArray[T](heapSize)
	pq.heap = arr
	for i := 0; i < heapSize; i++ {
		pq.heap.Add(elems[i])
	}
	for j := max(0, (heapSize/2)-1); j >= 0; j-- {
		pq.sink(j)
	}

	return pq
}

// Size returns the number of elements in the priority queue.
func (pq *BinaryHeapPQ[T]) Size() int {
	return pq.heap.Size()
}

// sink moves the element at index k down the tree until it satisfies the heap property.
func (pq *BinaryHeapPQ[T]) sink(k int) {
	heapSize := pq.heap.Size()
	for {
		left := 2*k + 1  //left node
		right := 2*k + 2 // right node

		smallest := left //left is smallest node assumption

		//find which is smaller
		// if right is smaller set smallest to right
		if right < heapSize && pq.less(right, left) {
			smallest = right
		}

		// check out of bound and not be able to sink K
		if left >= heapSize || pq.less(k, smallest) {
			break
		}

		// move down the tree
		pq.swap(smallest, k)
		k = smallest

	}
}

// swap swaps the elements at indices i and j in the heap.
func (pq *BinaryHeapPQ[T]) swap(i, j int) {
	elemI := pq.heap.Get(i)
	elemJ := pq.heap.Get(j)

	pq.heap.Set(i, elemJ)
	pq.heap.Set(j, elemI)
}

// less returns true if the element at index i is less than the element at index j in the heap.
func (pq *BinaryHeapPQ[T]) less(i, j int) bool {
	node1 := pq.heap.Get(i)
	node2 := pq.heap.Get(j)

	return node1 < node2
}

// Clear clears the heap
func (pq *BinaryHeapPQ[T]) Clear() {

}

// Peek returns the refrence of the element with the lowest
// priority in this pq. if pq is empty returns null
func (pq *BinaryHeapPQ[T]) Peek() *int {
	if pq.heap.IsEmpty() {
		return nil
	}

	val := pq.heap.Get(0)
	return &val

}

// Poll removeds the root of heap, O(log(n))
func (pq *BinaryHeapPQ[T]) Poll() *int {
	return pq.RemoveAt(0)
}

// RemoveAt removes node at particular index, O(log(n))
func (pq *BinaryHeapPQ[T]) RemoveAt(i int) *int {
	if pq.heap.IsEmpty() {
		return nil
	}

	indexOfLastElem := pq.Size() - 1
	removedData := pq.heap.Get(i)
	pq.swap(i, indexOfLastElem)

	pq.heap.RemoveAt(indexOfLastElem)

	// check if last elem was removed
	if i == indexOfLastElem {
		return &removedData
	}
	elem := pq.heap.Get(i)

	pq.sink(i)

	if reflect.DeepEqual(pq.heap.Get(i), elem) {
		pq.swim(i)
	}
	return &removedData
}

// remove removesa particular element in the heap. O(n)
// func (pq *BinaryHeapPQ[T]) remove(elem int) bool {
// 	for i := 0; i < pq.Size(); i++ {
// 		if reflect.DeepEqual(elem, pq.heap.Get(i)) {
// 			pq.RemoveAt(i)
// 			return true
// 		}
// 	}
// 	return false
// }

// swim bottom up node swim. O(log(n))
func (pq *BinaryHeapPQ[T]) swim(k int) {
	parent := (k - 1) / 2

	for k > 0 && pq.less(k, parent) {
		pq.swap(parent, k)
		k = parent

		// grab index of nexct parent
		parent = (k - 1) / 2
	}
}

// IsMinHeap recursive check that heap invariant is maitained
func (pq *BinaryHeapPQ[T]) IsMinHeap(k int) bool {
	heapSize := pq.Size()
	if k >= heapSize {
		return true
	}
	left := 2*k + 1
	right := 2*k + 2

	// Make sure that the current node k is less than
	// both of its children left, and right if they exist
	// return false otherwise to indicate an invalid heap

	if left < heapSize && !pq.less(k, left) {
		return false
	}
	if right < heapSize && !pq.less(k, right) {
		return false
	}

	return pq.IsMinHeap(left) && pq.IsMinHeap(right)

}

// Add adds a non null element to the PQ,, O(long(n))
func (pq *BinaryHeapPQ[T]) Add(elem int) {
	pq.heap.Add(elem)
	indexOfLastElem := pq.Size() - 1
	pq.swim(indexOfLastElem)
}

// Contains test if a elemen in heap, O(n)
func (pq *BinaryHeapPQ[T]) Contains(elem int) bool {
	for i := 0; i < pq.Size(); i++ {
		if reflect.DeepEqual(pq.heap.Get(i), elem) {
			return true
		}
	}
	return false
}

// ToSting string representaion of PQ
func (pq *BinaryHeapPQ[T]) ToString() string {
	return pq.heap.ToString()
}
