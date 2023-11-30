package singlylinkedlist

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	tests := []struct {
		name string
		want *SinglyLinkedList[int]
	}{
		{
			"TestNewSinglyLinkedList", &SinglyLinkedList[int]{0, nil, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSinglyLinkedList[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSinglyLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func get_set_Helper(size int) *SinglyLinkedList[int] {
	singlylinkedlist := NewSinglyLinkedList[int]()
	for i := 0; i < size; i++ {
		singlylinkedlist.Add(i)
	}

	return singlylinkedlist
}

func TestSinglyLinked_Clear(t *testing.T) {
	tests := []struct {
		name string
		s    *SinglyLinkedList[int]
		want int
	}{
		{
			"TestingSinglyLinkedList_Clear", get_set_Helper(10), 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Clear()
			if got := tt.s.Size(); got != tt.want {
				t.Errorf("SinglyLinkedList Size() = %v,  want %v", got, tt.want)
			}
		})
	}
}

func TestSinglyLinked_ToString(t *testing.T) {
	tests := []struct {
		name string
		s    *SinglyLinkedList[int]
		want string
	}{
		{
			"TestingSinglyLinked_ToString", get_set_Helper(10), "[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]",
		},
		{
			"TestSinglyLinkedList_ToString", NewSinglyLinkedList[int](), "[]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToString(); got != tt.want {
				t.Errorf("SinglyLinkedList.ToString() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestSinglyLinked_AddFirst(t *testing.T) {
	type args struct {
		elem int
	}

	tests := []struct {
		name string
		args args
		sl   *SinglyLinkedList[int]
		want int
	}{
		{
			"TestSinglyLinkedList_AddFirst", args{23}, get_set_Helper(10), 23,
		},
		{
			"TestSinglyLinkedList_AddFirst", args{10}, NewSinglyLinkedList[int](), 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sl.AddFirst(tt.args.elem); got.val != tt.want {
				t.Errorf("SinglyLinkedList.AddFirst() = %v, want %v", got.val, tt.want)
			}
		})
	}
}

func TestSinglyLinked_AddAt(t *testing.T) {
	type args struct {
		index int
		val   int
	}

	tests := []struct {
		name    string
		args    args
		sl      *SinglyLinkedList[int]
		wantErr error
	}{
		{
			"TestSinglyLinkedList_AddAt", args{0, 23}, get_set_Helper(10), nil,
		},
		{
			"TestSinglyLinkedList_AddAt", args{10, 23}, get_set_Helper(10), nil,
		},
		{
			"TestSinglyLinkedList_AddAt", args{21, 10}, get_set_Helper(10), errors.New("illegal index"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.sl.AddAt(tt.args.index, tt.args.val)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("SinglyLinkedList.AddAt() = %v, want %v", err, tt.wantErr)
			}
		})
	}

}

func TestSinglyLinked_Add(t *testing.T) {
	type args struct {
		elem int
	}

	tests := []struct {
		name string
		args args
		sl   *SinglyLinkedList[int]
		want int
	}{
		{
			"TestSinglyLinkedList_Add", args{23}, get_set_Helper(10), 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sl.Add(tt.args.elem)
			if got := tt.sl.Size(); got != tt.want {
				t.Errorf("SinglyLinkedList.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinglyLinked_AddRemove(t *testing.T) {

	tests := []struct {
		name string
		sl   *SinglyLinkedList[int]
		want error
	}{
		{
			"TestSinglyLinkedList_AddRemove", get_set_Helper(10), nil,
		},
		{
			"TestSinglyLinkedList_AddRemove", get_set_Helper(1), nil,
		},
		{
			"TestSinglyLinkedList_AddRemove", NewSinglyLinkedList[int](), errors.New("empty list"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.sl.Remove()
			// fmt.Println(tt.sl.ToString())
			if err != nil && err.Error() != tt.want.Error() {
				t.Errorf("SinglyLinkedList.Remove() = %v, want %v", err, tt.want)
			}
		})
	}
}
