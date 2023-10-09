package darray

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/exp/slices"
)

// DynamicArray is a struct that represents a dynamic array.
type DynamicArray struct {
	array    []int
	len      int
	capacity int
	size     int
}

// NewDynamicArray creates a new dynamic array with the given capacity.
// If no capacity is given, it defaults to 16.
func NewDynamicArray(args ...int) *DynamicArray {
	dynamicArray := &DynamicArray{
		array:    make([]int, 0, 16),
		len:      0,
		capacity: 16,
		size:     0,
	}

	if len(args) > 0 {
		if args[0] < 0 {
			log.Println("Illegal capacity defaulting to size 16")
		}
		args[0] = 16
		dynamicArray.capacity = args[0]
		dynamicArray.array = make([]int, 0, args[0])
	}

	return dynamicArray
}

// Get returns the element at the given index.
func (d *DynamicArray) Get(index int) int {
	return d.array[index]
}

// Set sets the element at the given index to the given value.
func (d *DynamicArray) Set(index int, elem int) {
	d.array[index] = elem
}

// IsEmpty returns true if the dynamic array is empty.
func (d *DynamicArray) IsEmpty() bool {
	return d.len == 0
}

// Size returns the number of elements in the dynamic array.
func (d *DynamicArray) Size() int {
	return d.len
}

// Add adds the given element to the dynamic array.
func (d *DynamicArray) Add(elem int) {
	if (d.len + 1) >= d.capacity {
		if d.capacity == 0 {
			d.capacity = 1
		} else {
			d.capacity *= 2
		}

		bigarray := make([]int, d.capacity)
		copy(bigarray, d.array)
		log.Println("Expanding capacity: ", d.capacity)
		d.array = bigarray

	}
	d.array = slices.Insert(d.array, d.len, elem)
	d.len += 1
}

// RemoveAt removes the element at the given index from the dynamic array.
func (d *DynamicArray) RemoveAt(index int) {
	d.array = append(d.array[:index], d.array[index+1:]...)
	d.len--
	d.capacity--
}

// Remove removes the given element from the dynamic array.
func (d *DynamicArray) Remove(elem int) error {
	for i := 0; i < d.len; i++ {
		if d.array[i] == elem {
			d.RemoveAt(i)
			return nil
		}
	}

	return fmt.Errorf("element %v not found", elem)
}

// ToString returns a string representation of the dynamic array.
func (d *DynamicArray) ToString() string {
	if d.len == 0 {
		return "[]"
	}
	var builder strings.Builder
	builder.WriteString("[")
	for i := 0; i < d.len-1; i++ {
		builder.WriteString(fmt.Sprintf("%v, ", d.array[i]))
	}
	builder.WriteString(fmt.Sprintf("%v]", d.array[d.len-1]))
	return builder.String()
}

// Reverse reverses the order of the elements in the dynamic array.
func (d *DynamicArray) Reverse() {
	for i := 0; i < d.len/2; i++ {
		tmp := d.array[i]
		d.array[i] = d.array[d.len-i-1]
		d.array[d.len-i-1] = tmp
	}
}

// Qsort sorts the dynamic array using the quicksort algorithm.
func (d *DynamicArray) Qsort() {
	quickSortHelper(d.array, 0, d.len-1)
}

func quickSortHelper(array []int, first int, last int) {
	if first < last {
		splitpoint := partition(array, first, last)

		quickSortHelper(array, first, splitpoint-1)
		quickSortHelper(array, splitpoint+1, last)
	}
}

func partition(array []int, first int, last int) int {
	pivot := array[first]

	leftmark := first + 1
	rightmark := last

	done := false

	for !done {
		for leftmark <= rightmark && array[leftmark] <= pivot {
			leftmark += 1
		}
		for array[rightmark] >= pivot && rightmark >= leftmark {
			rightmark -= 1
		}

		if rightmark < leftmark {
			done = true
		} else {
			temp := array[leftmark]
			array[leftmark] = array[rightmark]
			array[rightmark] = temp
		}
	}
	temp := array[first]
	array[first] = array[rightmark]
	array[rightmark] = temp

	return rightmark
}
