package binarysearchtree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewBST(t *testing.T) {
	tests := []struct {
		name string
		want *BST[int]
	}{
		{
			name: "TestNewBST with int",
			want: &BST[int]{},
		},
		{
			name: "TestNewBST with string",
			want: &BST[int]{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBST[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_Add(t *testing.T) {
	type args struct {
		elem int
	}
	tests := []struct {
		name string
		bst  *BST[int]
		args args
		want bool
	}{
		{
			name: "TestBST_Add with int",
			bst:  NewBST[int](),
			args: args{elem: 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.bst.Add(tt.args.elem); got != tt.want {
				t.Errorf("BST.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_ToString(t *testing.T) {
	tests := []struct {
		name string
		bst  *BST[int]
		want []string
	}{
		{
			name: "TestBST_ToString with int",
			bst:  createBST(),
			want: []string{"1", "4", "6", "7", "8", "10", "13", "14"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.bst.ToString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BST.ToString() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestNode_CompareTo(t *testing.T) {
	type args struct {
		value interface{}
	}
	var values int = 2
	tests := []struct {
		name string
		node *Node[int]
		args args
		want int
	}{
		{
			name: "TestNode_CompareTo with int",
			node: &Node[int]{value: values},
			args: args{value: 3},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tt.node.CompareTo(tt.args.value); got != tt.want {
				if err != nil {
					fmt.Println(err)
				}
				t.Errorf("Node.CompareTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

// helper function to create a BST for testing
func createBST() *BST[int] {
	bst := NewBST[int]()
	bst.Add(8)
	bst.Add(10)
	bst.Add(1)
	bst.Add(6)
	bst.Add(4)
	bst.Add(7)
	bst.Add(14)
	bst.Add(13)
	fmt.Println(bst.nodeCount)
	return bst
}

func TestBST_InorderTraversal(t *testing.T) {
	type args struct {
		node *Node[int]
	}
	tests := []struct {
		name string
		bst  *BST[int]
		args args
		want []int
	}{
		{
			name: "TestBST_InorderTraversal with int",
			bst:  createBST(),
			args: args{node: createBST().root},
			want: []int{1, 4, 6, 7, 8, 10, 13, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bst.InOrderTraversal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("type mismatch between %v and %v", reflect.TypeOf(got), reflect.TypeOf(tt.want))
				t.Errorf("BST.InorderTraversal() = %v want %v", got, tt.want)
			}
		})
	}

}
