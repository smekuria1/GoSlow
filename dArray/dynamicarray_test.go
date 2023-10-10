// Package darray provides an implementation of a dynamic array

// data structure.

package darray

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestNewDynamicArray(t *testing.T) {
	type args struct {
		args []int
	}
	tests := []struct {
		name string
		args args
		want *DynamicArray
	}{
		{"TestNewDynamicArray", args{[]int{}}, &DynamicArray{array: []int{}, len: 0, capacity: 16, size: 0}},
		{"TestNewDynamicArray", args{[]int{12}}, &DynamicArray{array: []int{}, len: 0, capacity: 12, size: 0}},
		{"TestNewDynamicArray", args{[]int{-1}}, &DynamicArray{array: []int{}, len: 0, capacity: 16, size: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDynamicArray(tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDynamicArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
func get_set_Helper(size int) *DynamicArray {
	dynamicArray := NewDynamicArray(size)
	for i := 0; i < 10; i++ {
		dynamicArray.Add(i)
	}
	return dynamicArray
}

func empty_Helper() *DynamicArray {
	dynamicArray := NewDynamicArray()
	return dynamicArray
}

func TestDynamicArray_Get(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		d    *DynamicArray
		args args
		want int
	}{
		// TODO: Add test cases.
		// create an array with 10 elements and insert 10 elements
		// get the 5th element and check if it is the same as the 5th element
		// in the array
		{"TestDynamicArray_Get", get_set_Helper(10), args{5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Get(tt.args.index); got != tt.want {
				t.Errorf("DynamicArray.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamicArray_Set(t *testing.T) {
	type args struct {
		index int
		elem  int
	}
	tests := []struct {
		name string
		d    *DynamicArray
		args args
	}{
		// TODO: Add test cases.
		// create an array with 10 elements and insert 10 elements
		// set the 5th element to 10 and check if it is the same as the 5th element
		// in the array
		{"TestDynamicArray_Set", get_set_Helper(10), args{5, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Set(tt.args.index, tt.args.elem)
			if got := tt.d.Get(tt.args.index); got != tt.args.elem {
				t.Errorf("DynamicArray.Get() = %v, want %v", got, tt.args.elem)
			}

		})
	}
}

func TestDynamicArray_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		d    *DynamicArray
		want bool
	}{
		// TODO: Add test cases.
		// create an array with 10 elements and insert 10 elements
		// check if the array is empty
		{"TestDynamicArray_IsEmpty", get_set_Helper(10), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.IsEmpty(); got != tt.want {
				t.Errorf("DynamicArray.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamicArray_Size(t *testing.T) {
	tests := []struct {
		name string
		d    *DynamicArray
		want int
	}{
		// TODO: Add test cases.
		// create an array with 10 elements and insert 10 elements
		// check if the size of the array is 10
		{"TestDynamicArray_Size", get_set_Helper(10), 10},
		// initialize with 0 capacity
		{"TestDynamicArray_Size", get_set_Helper(0), 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Size(); got != tt.want {
				t.Errorf("DynamicArray.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamicArray_Add(t *testing.T) {
	type args struct {
		elem int
	}
	tests := []struct {
		name string
		d    *DynamicArray
		args args
	}{
		// TODO: Add test cases.
		// create an array with 10 elements and insert 10 elements
		// add an element to the array and check if the size of the array is 11
		{"TestDynamicArray_Add", get_set_Helper(10), args{11}},
		// initialize with 0 capacity
		{"TestDynamicArray_Add", get_set_Helper(0), args{11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Add(tt.args.elem)
			if got := tt.d.Size(); got != 11 {
				t.Errorf("DynamicArray.Size() = %v, want %v", got, 11)
			}

		})
	}
}

func TestDynamicArray_RemoveAt(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		d    *DynamicArray
		args args
	}{
		// TODO: Add test cases.
		// create an array with 10 elements and insert 10 elements
		// remove the 5th element and check if the size of the array is 9
		{"TestDynamicArray_RemoveAt", get_set_Helper(10), args{5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.RemoveAt(tt.args.index)
			if got := tt.d.Size(); got != 9 {
				t.Errorf("DynamicArray.Size() = %v, want %v", got, 9)
			}
		})
	}
}

func TestDynamicArray_Remove(t *testing.T) {
	type args struct {
		elem int
	}
	tests := []struct {
		name    string
		d       *DynamicArray
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		// create an array with 10 elements and insert 10 elements
		// remove the 5th element and check if the size of the array is 9
		{"TestDynamicArray_Remove", get_set_Helper(10), args{5}, false},
		{"TestDynamicArray_Remove", get_set_Helper(10), args{11}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Remove(tt.args.elem); (err != nil) != tt.wantErr {
				t.Errorf("DynamicArray.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDynamicArray_ToString(t *testing.T) {
	tests := []struct {
		name string
		d    *DynamicArray
		want string
	}{
		// TODO: Add test cases.
		// create an array with 10 elements and insert 10 elements
		// check if the string representation of the array is correct
		{"TestDynamicArray_ToString", get_set_Helper(10), "[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]"},
		// empty array
		{"TestDynamicArray_ToString", empty_Helper(), "[]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.ToString(); got != tt.want {
				t.Errorf("DynamicArray.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamicArray_Reverse(t *testing.T) {
	tests := []struct {
		name string
		d    *DynamicArray
	}{
		// TODO: Add test cases.
		// create an array with 10 elements and insert 10 elements
		// reverse the array and check if the string representation of the array is correct
		{"TestDynamicArray_Reverse", get_set_Helper(10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Reverse()
			if got := tt.d.ToString(); got != "[9, 8, 7, 6, 5, 4, 3, 2, 1, 0]" {
				t.Errorf("DynamicArray.ToString() = %v, want %v", got, "[9, 8, 7, 6, 5, 4, 3, 2, 1, 0]")
			}
		})
	}
}

func TestDynamicArray_Qsort(t *testing.T) {
	tests := []struct {
		name string
		d    *DynamicArray
	}{
		// TODO: Add test cases.
		// create an array with 10 elements and insert 10 random elements from 1 to 10
		// sort the array and check if the string representation of the array is correct
		{"TestDynamicArray_Qsort", sortHelper(10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Qsort()
			if got := tt.d.ToString(); got != "[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]" {
				t.Errorf("DynamicArray.ToString() = %v, want %v", got, "[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]")
			}

		})
	}
}

// random generated nums for sorting each number is between 1 and 10
func sortHelper(size int) *DynamicArray {
	dynamicArray := NewDynamicArray(size)
	for i := 0; i < size; i++ {
		dynamicArray.Add(rand.Intn(10-1) + 1)
	}
	return dynamicArray
}

func TestDynamicArray_Contains(t *testing.T) {
	type args struct {
		elem int
	}
	tests := []struct {
		name string
		d    *DynamicArray
		args args
		want bool
	}{
		// TODO: Add test cases.
		// create an array with 10 elements and insert 10 elements
		// check if the array contains the 5th element
		{"TestDynamicArray_Contains", get_set_Helper(10), args{5}, true},
		{"TestDynamicArray_Contains", get_set_Helper(10), args{11}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Contains(tt.args.elem); got != tt.want {
				t.Errorf("DynamicArray.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
