// Package binarysearchtree implements a binary search tree data structure.
package binarysearchtree

import (
	"fmt"
	"log"
	"reflect"

	"github.com/smekuria1/GoSlow/queue"
)

// Node represents a node in the binary search tree.
type Node[T comparable] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

// BST represents a binary search tree.
type BST[T comparable] struct {
	root      *Node[T]
	nodeCount int
}

// NewBST creates a new binary search tree.
func NewBST[T comparable]() *BST[T] {
	return &BST[T]{root: nil, nodeCount: 0}
}

// IsEmpty checks if the binary tree is empty
func (bst *BST[T]) IsEmpty() bool {
	return bst.Size() == 0
}

// Size get the number of nodes in the binary tree
func (bst *BST[T]) Size() int {
	return bst.nodeCount
}

// Add an element to this binary tree. Returns true
// if we successfully perform an insertion
func (bst *BST[T]) Add(elem T) bool {
	// check if the value already exists in this
	//binary tree, if it does ignore
	if bst.Contains(elem) {
		return false
	} else {
		bst.root = bst.add(bst.root, elem)
		bst.nodeCount++
		return true
	}

}

// private method to recursively add a value in the binary tree
func (bst *BST[T]) add(node *Node[T], elem T) *Node[T] {
	// base case: found a leaf node
	if node == nil {
		node = &Node[T]{value: elem, left: nil, right: nil}
	} else {
		// place lower elements values in the left subtree
		cmp, err := node.CompareTo(elem)
		if err != nil {
			return node
		}

		if cmp > 0 {
			node.left = bst.add(node.left, elem)
		} else {
			node.right = bst.add(node.right, elem)
		}
	}

	return node
}

// Contains checks whether the given element is in the tree
func (bst *BST[T]) Contains(elem T) bool {
	return bst.contains(bst.root, elem)
}

// private recursive method to find an element in the tree
func (bst *BST[T]) contains(node *Node[T], elem T) bool {
	if node == nil {
		return false
	}

	cmp, err := node.CompareTo(elem)
	if err != nil {
		log.Fatal(err)
		return false
	}

	if cmp > 0 {
		return bst.contains(node.left, elem)
	} else if cmp < 0 {
		return bst.contains(node.right, elem)
	} else {
		return true
	}
}

// private compare method to compare two elements
func (n *Node[T]) CompareTo(value interface{}) (int, error) {
	v1 := reflect.ValueOf(n.value)
	v2 := reflect.ValueOf(value)

	if v1.Type() != v2.Type() {
		return 0, fmt.Errorf("type mismatch between %v and %v", v1.Type(), v2.Type())
	}

	switch v1.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v1.Int() < v2.Int() {
			return -1, nil
		} else if v1.Int() > v2.Int() {
			return 1, nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v1.Uint() < v2.Uint() {
			return -1, nil
		} else if v1.Uint() > v2.Uint() {
			return 1, nil
		}
	case reflect.Float32, reflect.Float64:
		if v1.Float() < v2.Float() {
			return -1, nil
		} else if v1.Float() > v2.Float() {
			return 1, nil
		}
	case reflect.String:
		if v1.String() < v2.String() {
			return -1, nil
		} else if v1.String() > v2.String() {
			return 1, nil
		}
	default:
		return 0, fmt.Errorf("type not supported for comparison")
	}

	return 0, nil
}

// Remove removes an element from this binary tree if it exists, O(n)
func (bst *BST[T]) Remove(elem T) bool {
	if bst.Contains(elem) {
		bst.root = bst.remove(bst.root, elem)
		bst.nodeCount--
		return true
	}
	return false
}

// private method to remove a node from the tree
func (bst *BST[T]) remove(node *Node[T], elem T) *Node[T] {
	if node == nil {
		return nil
	}

	cmp, err := node.CompareTo(elem)
	if err != nil {
		return node
	}

	if cmp < 0 {
		node.left = bst.remove(node.left, elem)
	} else if cmp > 0 {
		node.right = bst.remove(node.right, elem)
	} else {
		if node.left == nil {
			rightChild := node.right
			//node.value = nil
			node = nil
			return rightChild
		} else if node.right == nil {
			leftChild := node.left
			//node.value = nil
			node = nil
			return leftChild
		} else {
			// find the minimum element in the right subtree
			tmp := bst.findMin(node.right)
			node.value = tmp.value
			node.right = bst.remove(node.right, tmp.value)
		}
	}

	return node
}

// FindMax finds the maximum value in the binary tree
func (bst *BST[T]) FindMax() (*T, error) {
	if bst.IsEmpty() {
		return nil, fmt.Errorf("tree is empty")
	}
	return &bst.findMax(bst.root).value, nil
}

// Helper method to find the rightmost node (which has the largest value)
func (bst *BST[T]) findMax(node *Node[T]) *Node[T] {
	for node.right != nil {
		node = node.right
	}
	return node
}

// Helper method to find the leftmost node (which has the smallest value)
func (bst *BST[T]) findMin(node *Node[T]) *Node[T] {
	for node.left != nil {
		node = node.left
	}
	return node
}

// FindMin finds the minimum value in the binary tree
func (bst *BST[T]) FindMin() (*T, error) {
	if bst.IsEmpty() {
		return nil, fmt.Errorf("tree is empty")
	}
	return &bst.findMin(bst.root).value, nil
}

// private height method
// func (bst *BST[T]) height(node *Node[T]) int {
// 	if node == nil {
// 		return 0
// 	}
// 	return max(bst.height(node.left), bst.height(node.right)) + 1
// }

// ToStrArray returns a string array representation of the binary tree
func (bst *BST[T]) ToString() []string {
	return bst.toStrArray(bst.root)
}

// private method to convert the binary tree to a string array
func (bst *BST[T]) toStrArray(node *Node[T]) []string {
	if node == nil {
		return []string{}
	}
	fmt.Println("node: ", node.value)
	fmt.Println("left: ", node.left)
	fmt.Println("right: ", node.right)
	return append(append(append([]string{}, bst.toStrArray(node.left)...), fmt.Sprintf("%v", node.value)), bst.toStrArray(node.right)...)
}

// InOrderTraversal traverses the binary tree in order
func (bst *BST[T]) InOrderTraversal() []T {
	return bst.inOrderTraversal(bst.root)
}

// private recursive method to traverse the binary tree in order
func (bst *BST[T]) inOrderTraversal(node *Node[T]) []T {
	if node == nil {
		return []T{}
	}

	var result []T
	result = append(result, bst.inOrderTraversal(node.left)...)
	result = append(result, node.value)
	result = append(result, bst.inOrderTraversal(node.right)...)
	return result

}

// PreOrderTraversal traverses the binary tree in pre order
func (bst *BST[T]) PreOrderTraversal() []T {
	return bst.preOrderTraversal(bst.root)
}

// private recursive method to traverse the binary tree in pre order
func (bst *BST[T]) preOrderTraversal(node *Node[T]) []T {
	if node == nil {
		return []T{}
	}

	var result []T
	result = append(result, node.value)
	result = append(result, bst.preOrderTraversal(node.left)...)
	result = append(result, bst.preOrderTraversal(node.right)...)
	return result

}

// PostOrderTraversal traverses the binary tree in post order
func (bst *BST[T]) PostOrderTraversal() []T {
	return bst.postOrderTraversal(bst.root)
}

// private recursive method to traverse the binary tree in post order
func (bst *BST[T]) postOrderTraversal(node *Node[T]) []T {
	if node == nil {
		return []T{}
	}

	var result []T
	result = append(result, bst.postOrderTraversal(node.left)...)
	result = append(result, bst.postOrderTraversal(node.right)...)
	result = append(result, node.value)
	return result

}

// LevelOrderTraversal traverses the binary tree in level order
func (bst *BST[T]) LevelOrderTraversal() []T {
	return bst.levelOrderTraversal(bst.root)
}

// private recursive method to traverse the binary tree in level order
func (bst *BST[T]) levelOrderTraversal(node *Node[T]) []T {
	if node == nil {
		return []T{}
	}

	var result []T
	queue := queue.NewQueue[*Node[T]]()
	queue.Enqueue(node)

	for !queue.IsEmpty() {
		node := queue.Dequeue()
		result = append(result, node.value)
		if node.left != nil {
			queue.Enqueue(node.left)
		}
		if node.right != nil {
			queue.Enqueue(node.right)
		}
	}

	return result

}
