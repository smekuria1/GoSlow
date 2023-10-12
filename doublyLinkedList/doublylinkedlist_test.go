// Package doublyLinkedList provides an implementation of a doubly linked list

// data structure.

package doublyLinkedList

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestNewDoublyLinkedList(t *testing.T) {
	tests := []struct {
		name string
		want *DoublyLinkedList[int]
	}{
		// TODO: Add test cases.
		{"TestNewDoublyLinkedListInt", &DoublyLinkedList[int]{size: 0, head: nil, tail: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoublyLinkedList[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDoublyLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestNewDoublyLinkedListString(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want *DoublyLinkedList[string]
// 	}{}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewDoublyLinkedList[string](); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewDoublyLinkedList() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestNewNode(t *testing.T) {
	type args struct {
		val  int
		prev *Node[int]
		next *Node[int]
	}
	tests := []struct {
		name string
		args args
		want *Node[int]
	}{
		// TODO: Add test cases.
		{"TestNewNode", args{val: 1, prev: nil, next: nil}, &Node[int]{next: nil, prev: nil, val: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.val, tt.args.prev, tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to create a doubly linked list with 10 elements
func get_set_Helper(size int) *DoublyLinkedList[int] {
	doublyLinkedList := NewDoublyLinkedList[int]()
	for i := 0; i < 10; i++ {
		doublyLinkedList.Add(i)
	}
	return doublyLinkedList
}

func TestDoublyLinkedList_Size(t *testing.T) {
	tests := []struct {
		name string
		d    *DoublyLinkedList[int]
		want int
	}{
		//TODO: Add test cases
		{testNameHelper("Size"), &DoublyLinkedList[int]{size: 0, head: nil, tail: nil}, 0},
		{testNameHelper("Size"), get_set_Helper(10), 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Size(); got != tt.want {
				t.Errorf("%s %v, want %v", testErrorNameHelper("Size"), got, tt.want)
			}
		})
	}

}

func TestDoublyLinkedList_Clear(t *testing.T) {
	tests := []struct {
		name string
		d    *DoublyLinkedList[int]
	}{
		//TODO: Add test cases
		{testErrorNameHelper("Clear"), get_set_Helper(10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Clear()
			if tt.d.Size() != 0 {
				t.Errorf("%s %v, want %v", testErrorNameHelper("Clear"), tt.d.Size(), 0)
			}
		})
	}
}

func TestDoublyLinkedList_Add(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		d    *DoublyLinkedList[int]
		args args
	}{
		//TODO: Add test cases
		{testNameHelper("Add"), get_set_Helper(10), args{val: 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Add(tt.args.val)
			if tt.d.Size() != 11 {
				t.Errorf("%s %v, want %v", testErrorNameHelper("Add"), tt.d.Size(), 11)

			}
		})
	}
}

func TestDoublyLinkedList_AddFirst(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		d    *DoublyLinkedList[int]
		args args
	}{
		//TODO: Add test cases
		{testNameHelper("AddFirst"), get_set_Helper(10), args{val: 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.AddFirst(tt.args.val)
			if tt.d.Size() != 11 {
				t.Errorf("%s %v, want %v", testErrorNameHelper("AddFirst"), tt.d.Size(), 11)
			}
		})
	}
}

func TestDoublyLinkedList_AddAt(t *testing.T) {
	type args struct {
		index int
		val   int
	}
	tests := []struct {
		name string
		d    *DoublyLinkedList[int]
		args args
	}{
		//TODO: Add test cases
		{testNameHelper("AddAt"), get_set_Helper(10), args{index: 5, val: 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.AddAt(tt.args.index, tt.args.val)
			if tt.d.IndexOf(tt.args.val) != tt.args.index {
				t.Errorf("%s = %v, want %v", testErrorNameHelper("AddAt"), tt.d.IndexOf(tt.args.val), tt.args.index)
			}
		})
	}
}

func TestDoublyLinkedList_ToString(t *testing.T) {
	tests := []struct {
		name string
		d    *DoublyLinkedList[int]
		want string
	}{
		{testNameHelper("ToString"), get_set_Helper(10), "[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.ToString(); got != tt.want {
				t.Errorf("%s %v, want %v", testErrorNameHelper("ToString"), got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_PeekFirst(t *testing.T) {
	tests := []struct {
		name string
		d    *DoublyLinkedList[int]
		want int
	}{
		{testNameHelper("PeekFirst"), get_set_Helper(10), 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.PeekFirst(); got != tt.want {
				t.Errorf("%s %v, want %v", testErrorNameHelper("PeekFirst"), got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_PeekLast(t *testing.T) {
	tests := []struct {
		name string
		d    *DoublyLinkedList[int]
		want int
	}{
		{testNameHelper("PeekLast"), get_set_Helper(10), 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.PeekLast(); got != tt.want {
				t.Errorf("%s %v, want %v", testErrorNameHelper("PeekLast"), got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_RemoveLast(t *testing.T) {
	tests := []struct {
		name string
		d    *DoublyLinkedList[int]
		want int
	}{
		{testNameHelper("RemoveLast"), get_set_Helper(10), 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.RemoveLast(); got != tt.want {
				t.Errorf("%s %v, want %v", testErrorNameHelper("RemoveLast"), got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_RemoveFirst(t *testing.T) {
	tests := []struct {
		name string
		d    *DoublyLinkedList[int]
		want int
	}{
		{testNameHelper("RemoveFirst"), get_set_Helper(10), 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.RemoveFirst(); got != tt.want {
				t.Errorf("%s %v, want %v", testErrorNameHelper("RemoveFirst"), got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_RemoveAt(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		d    *DoublyLinkedList[int]
		args args
		want int
	}{
		{testNameHelper("RemoveAt"), get_set_Helper(10), args{5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.RemoveAt(tt.args.index); got != tt.want {
				t.Errorf("%s %v, want %v", testErrorNameHelper("RemoveAt"), got, tt.want)
			}
		})
	}
}
func testErrorNameHelper(methname string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("DoublyLinkedList.%s()=", methname))
	return builder.String()

}

func testNameHelper(testname string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("TestingDoublyLinkedList_%s", testname))
	return builder.String()
}

func TestDoublyLinkedList_RemoveVal(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		d    *DoublyLinkedList[int]
		args args
		want bool
	}{
		{testNameHelper("RemoveVal"), get_set_Helper(10), args{7}, true},
		{testNameHelper("RemoveValFalse"), get_set_Helper(10), args{22}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.RemoveVal(tt.args.val); got != tt.want {
				t.Errorf("%s %v, want %v", testErrorNameHelper("RemoveVal"), got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_IndexOf(t *testing.T) {
	type args struct {
		elem int
	}
	tests := []struct {
		name string
		d    *DoublyLinkedList[int]
		args args
		want int
	}{
		{testNameHelper("IndexOf"), get_set_Helper(10), args{3}, 3},
		{testErrorNameHelper("IndexOfNotFound"), get_set_Helper(10), args{100}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.IndexOf(tt.args.elem); got != tt.want {
				t.Errorf("%s %v, want %v", testErrorNameHelper("IndexOf"), got, tt.want)
			}
		})
	}
}
