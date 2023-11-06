// Package priorityqueue implements a priority queue.

// uses my dynamicArray package for the heap

package priorityqueue

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/smekuria1/GoSlow/darray"
)

func TestNewBinaryHeapPQ(t *testing.T) {
	type args struct {
		args []int
	}

	type T int
	tests := []struct {
		name string
		args args
		want *BinaryHeapPQ[T]
	}{
		{
			name: "TestNewBinaryHeapPQ",
			args: args{
				args: []int{10},
			},
			want: &BinaryHeapPQ[T]{
				heap: darray.NewDynamicArray[T](10),
			},
		},
		{
			name: "TestNewBinaryHeapPQ",
			args: args{
				args: []int{},
			},
			want: &BinaryHeapPQ[T]{
				heap: darray.NewDynamicArray[T](),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinaryHeapPQ[T](tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinaryHeapPQ() = %v, want %v", got, tt.want)
			}
		})
	}
}

// helper function make a heap from an array
func makeHeap[T comparable](elems []int) *BinaryHeapPQ[T] {
	heapSize := len(elems)
	arr := darray.NewDynamicArray[T](heapSize)
	pq := &BinaryHeapPQ[T]{
		heap: arr,
	}
	for i := 0; i < heapSize; i++ {
		pq.heap.Add(elems[i])
	}
	for j := max(0, (heapSize/2)-1); j >= 0; j-- {
		pq.sink(j)
	}
	fmt.Println("Heap: ", pq.ToString())
	return pq
}

func TestBinaryHeapPQSize(t *testing.T) {
	tests := []struct {
		name string
		pq   *BinaryHeapPQ[int]
		want int
	}{
		{
			name: "TestBinaryHeapPQSize",
			pq:   makeHeap[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			want: 10,
		},
		{
			name: "TestBinaryHeapPQSize",
			pq:   makeHeap[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}),
			want: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Size(); got != tt.want {
				t.Errorf("BinaryHeapPQ.Size() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestBinaryHeapPQPeek(t *testing.T) {
	tests := []struct {
		name string
		pq   *BinaryHeapPQ[int]
		want int
	}{
		{
			name: "TestBinaryHeapPQPeek",
			pq:   makeHeap[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			want: 1,
		},
		{
			name: "TestBinaryHeapPQPeek",
			pq:   makeHeap[int]([]int{}),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.pq.Peek()
			if got == nil {
				got = new(int)
			}
			if *got != tt.want {
				t.Errorf("BinaryHeapPQ.Peek() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestBinaryHeapPQToString(t *testing.T) {
	tests := []struct {
		name string
		pq   *BinaryHeapPQ[int]
		want string
	}{
		{
			name: "TestBinaryHeapPQToString",
			pq:   makeHeap[int]([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}),
			want: "[1, 2, 3, 4, 7, 9, 10, 14, 8, 16]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.ToString(); got != tt.want {
				t.Errorf("BinaryHeapPQ.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryHeapPQIsMinHeap(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		pq   *BinaryHeapPQ[int]
		args args
		want bool
	}{
		{
			name: "TestBinaryHeapPQIsMinHeap",
			pq:   makeHeap[int]([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}),
			args: args{
				i: 0,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.IsMinHeap(tt.args.i); got != tt.want {
				t.Errorf("BinaryHeapPQ.IsMinHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryHeapPQAdd(t *testing.T) {
	type args struct {
		elem int
	}
	tests := []struct {
		name string
		pq   *BinaryHeapPQ[int]
		args args
		want string
	}{
		{
			name: "TestBinaryHeapPQAdd",
			pq:   makeHeap[int]([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}),
			args: args{
				elem: 5,
			},
			want: "[1, 2, 3, 4, 5, 9, 10, 14, 8, 16, 7]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.Add(tt.args.elem)
			if got := tt.pq.ToString(); got != tt.want {
				t.Errorf("BinaryHeapPQ.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryHeapPQContains(t *testing.T) {
	type args struct {
		elem int
	}
	tests := []struct {
		name string
		pq   *BinaryHeapPQ[int]
		args args
		want bool
	}{
		{
			name: "TestBinaryHeapPQContains",
			pq:   makeHeap[int]([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}),
			args: args{
				elem: 5,
			},
			want: false,
		},
		{
			name: "TestBinaryHeapPQContains",
			pq:   makeHeap[int]([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}),
			args: args{
				elem: 16,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Contains(tt.args.elem); got != tt.want {
				t.Errorf("BinaryHeapPQ.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryHeapPQPoll(t *testing.T) {
	tests := []struct {
		name string
		pq   *BinaryHeapPQ[int]
		want int
	}{
		{
			name: "TestBinaryHeapPQPoll",
			pq:   makeHeap[int]([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}),
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Poll(); *got != tt.want {
				t.Errorf("BinaryHeapPQ.Poll() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestBinaryHeapPQRemoveAt(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		pq   *BinaryHeapPQ[int]
		args args
		want int
	}{
		{
			name: "TestBinaryHeapPQRemoveAt",
			pq:   makeHeap[int]([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}),
			args: args{
				i: 5,
			},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.RemoveAt(tt.args.i); *got != tt.want && tt.pq.IsMinHeap(0) {
				t.Errorf("BinaryHeapPQ.RemoveAt() = %v, want %v", *got, tt.want)
			}
		})
	}
}
