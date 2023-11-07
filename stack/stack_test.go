package stack

import (
	"reflect"
	"testing"

	"github.com/smekuria1/GoSlow/doublyLinkedList"
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		want *Stack[int]
	}{
		{
			name: "TestNewStack",
			want: &Stack[int]{
				list: doublyLinkedList.NewDoublyLinkedList[int](),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

// helper function to create a stack with 10 elements
func stack_helper() *Stack[int] {
	stack := NewStack[int]()
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	return stack
}

func TestStack_Push(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		s    *Stack[int]
		args args
		want int
	}{
		{
			name: "TestStack_Push",
			s:    stack_helper(),
			args: args{10},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.val)
			if got := tt.s.Size(); got != tt.want {
				t.Errorf("Stack.Push() = %v, want %v", got, tt.want)
			}
		})

	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack[int]
		want int
	}{
		{
			name: "TestStack_Pop",
			s:    stack_helper(),
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Pop(); got != tt.want {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack[int]
		want int
	}{
		{
			name: "TestStack_Peek",
			s:    stack_helper(),
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Peek(); got != tt.want {
				t.Errorf("Stack.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Size(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack[int]
		want int
	}{
		{
			name: "TestStack_Size",
			s: &Stack[int]{
				list: doublyLinkedList.NewDoublyLinkedList[int](),
			},
			want: 0,
		},
		{
			name: "TestStack_Size",
			s:    stack_helper(),
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Size(); got != tt.want {
				t.Errorf("Stack.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack[int]
		want bool
	}{
		{
			name: "TestStack_IsEmpty",
			s: &Stack[int]{
				list: doublyLinkedList.NewDoublyLinkedList[int](),
			},
			want: true,
		},
		{
			name: "TestStack_IsEmpty",
			s:    stack_helper(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsEmpty(); got != tt.want {
				t.Errorf("Stack.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Clear(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack[int]
		want int
	}{
		{
			name: "TestStack_Clear",
			s:    stack_helper(),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Clear()
			if got := tt.s.Size(); got != tt.want {
				t.Errorf("Stack.Clear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_ToString(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack[int]
		want string
	}{
		{
			name: "TestStack_ToString",
			s:    stack_helper(),
			want: "[9, 8, 7, 6, 5, 4, 3, 2, 1, 0]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToString(); got != tt.want {
				t.Errorf("Stack.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Contains(t *testing.T) {
	type args struct {
		val int
	}

	tests := []struct {
		name string
		s    *Stack[int]
		args args
		want bool
	}{
		{
			name: "TestStack_Contains",
			s:    stack_helper(),
			args: args{10},
			want: false,
		},
		{
			name: "TestStack_Contains",
			s:    stack_helper(),
			args: args{9},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.val); got != tt.want {
				t.Errorf("Stack.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
