package queue

import (
	"reflect"
	"testing"

	"github.com/smekuria1/GoSlow/doublyLinkedList"
)

func TestNewQueue(t *testing.T) {
	tests := []struct {
		name string
		want *Queue[int]
	}{
		{
			name: "TestNewQueue",
			want: &Queue[int]{
				list: doublyLinkedList.NewDoublyLinkedList[int](),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

// helper function to create a queue with 10 elements
func queue_helper() *Queue[int] {
	queue := NewQueue[int]()
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
	}
	return queue
}

func TestQueue_Size(t *testing.T) {
	tests := []struct {
		name string
		q    *Queue[int]
		want int
	}{
		{
			name: "TestQueue_Size",
			q: &Queue[int]{
				list: doublyLinkedList.NewDoublyLinkedList[int](),
			},
			want: 0,
		},
		{
			name: "TestQueue_Size",
			q:    queue_helper(),
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue[int]{
				list: tt.q.list,
			}
			if got := q.Size(); got != tt.want {
				t.Errorf("Queue.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		q    *Queue[int]
		want bool
	}{
		{
			name: "TestQueue_IsEmpty",
			q: &Queue[int]{
				list: doublyLinkedList.NewDoublyLinkedList[int](),
			},
			want: true,
		},
		{
			name: "TestQueue_IsEmpty",
			q:    queue_helper(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue[int]{
				list: tt.q.list,
			}
			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("Queue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Clear(t *testing.T) {
	tests := []struct {
		name string
		q    *Queue[int]
	}{
		{
			name: "TestQueue_Clear",
			q:    queue_helper(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue[int]{
				list: tt.q.list,
			}
			q.Clear()
			if !q.IsEmpty() {
				t.Errorf("Queue.Clear() = %v, want %v", q.IsEmpty(), true)
			}
		})
	}
}

func TestQueue_Enqueue(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		q    *Queue[int]
		args args
	}{
		{
			name: "TestQueue_Enqueue",
			q:    queue_helper(),
			args: args{10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue[int]{
				list: tt.q.list,
			}
			q.Enqueue(tt.args.val)
			if got := q.Size(); got != 11 {
				t.Errorf("Queue.Enqueue() = %v, want %v", got, 11)
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name string
		q    *Queue[int]
		want int
	}{
		{
			name: "TestQueue_Dequeue",
			q:    queue_helper(),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue[int]{
				list: tt.q.list,
			}
			if got := q.Dequeue(); got != tt.want {
				t.Errorf("Queue.Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	tests := []struct {
		name string
		q    *Queue[int]
		want int
	}{
		{
			name: "TestQueue_Peek",
			q:    queue_helper(),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue[int]{
				list: tt.q.list,
			}
			if got := q.Peek(); got != tt.want {
				t.Errorf("Queue.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_ToString(t *testing.T) {
	tests := []struct {
		name string
		q    *Queue[int]
		want string
	}{
		{
			name: "TestQueue_ToString",
			q:    queue_helper(),
			want: "[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue[int]{
				list: tt.q.list,
			}
			if got := q.ToString(); got != tt.want {
				t.Errorf("Queue.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Contains(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		q    *Queue[int]
		args args
		want bool
	}{
		{
			name: "TestQueue_Contains",
			q:    queue_helper(),
			args: args{10},
			want: false,
		},
		{
			name: "TestQueue_Contains",
			q:    queue_helper(),
			args: args{0},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue[int]{
				list: tt.q.list,
			}
			if got := q.Contains(tt.args.val); got != tt.want {
				t.Errorf("Queue.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
