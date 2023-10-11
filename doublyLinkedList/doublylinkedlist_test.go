// Package doublyLinkedList provides an implementation of a doubly linked list

// data structure.

package doublyLinkedList

import (
	"reflect"
	"testing"
)

func TestNewDoublyLinkedListInt(t *testing.T) {
	tests := []struct {
		name string
		want *DoublyLinkedList[int]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoublyLinkedList[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDoublyLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDoublyLinkedListString(t *testing.T) {
	tests := []struct {
		name string
		want *DoublyLinkedList[string]
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoublyLinkedList[string](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDoublyLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
