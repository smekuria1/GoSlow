// Copyright 2023 Solomon Mekuria.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// GoSlow is combination of multiple packages that each represent popular
// algorithms and data structures. The goal of this project is to provide
// a simple and easy to use library for learning and teaching purposes.
// The library is not intended to be used in production code.
package main

import (
	"fmt"
	"math/rand"

	"github.com/smekuria1/GoSlow/binarysearchtree"
	"github.com/smekuria1/GoSlow/darray"
	"github.com/smekuria1/GoSlow/doublyLinkedList"
	"github.com/smekuria1/GoSlow/hashtable"
	"github.com/smekuria1/GoSlow/priorityqueue"
	"github.com/smekuria1/GoSlow/queue"
	"github.com/smekuria1/GoSlow/stack"
	"github.com/smekuria1/GoSlow/unionfind"
)

func main() {
	//Example of using the binarysearchtree package
	bst := binarysearchtree.NewBST[int]()
	// add random numbers generated from 1 to 100
	for i := 0; i < 10; i++ {
		bst.Add(rand.Intn(100))
	}
	// print the tree
	fmt.Println("Binary Search Tree")
	bst.ToString()

	//Example of using the darray package
	dynamicarray := darray.NewDynamicArray[int]()
	// add random numbers generated from 1 to 100
	for i := 0; i < 10; i++ {
		dynamicarray.Add(rand.Intn(100))
	}
	// print the array
	fmt.Println("Dynamic Array")
	fmt.Println(dynamicarray.ToString())

	//Example of using doublylinkedlist package
	dll := doublyLinkedList.NewDoublyLinkedList[int]()
	// add random numbers generated from 1 to 100
	for i := 0; i < 10; i++ {
		dll.Add(rand.Intn(100))
	}
	// print the list
	fmt.Println("Doubly Linked List")
	fmt.Println(dll.ToString())

	// Example of using the hashtable package
	ht := hashtable.NewHashTable[int, int](10, 0.75)
	// add random numbers generated from 1 to 100
	for i := 0; i < 10; i++ {
		ht.Insert(i, rand.Intn(100))
	}
	// print the hashtable
	fmt.Println("Hash Table")
	fmt.Println(ht.ToString())

	// Example of using priorityqueue package
	pq := priorityqueue.NewBinaryHeapPQ[int]()
	// add random numbers generated from 1 to 100
	for i := 0; i < 10; i++ {
		pq.Add(rand.Intn(100))
	}
	// print the priorityqueue
	fmt.Println("Priority Queue")
	fmt.Println(pq.ToString())

	// Example of using queue package
	q := queue.NewQueue[int]()
	// add random numbers generated from 1 to 100
	for i := 0; i < 10; i++ {
		q.Enqueue(rand.Intn(100))
	}
	// print the queue
	fmt.Println("Queue")
	fmt.Println(q.ToString())

	// Example of using stack package
	s := stack.NewStack[int]()
	// add random numbers generated from 1 to 100
	for i := 0; i < 10; i++ {
		s.Push(rand.Intn(100))
	}
	// print the stack
	fmt.Println("Stack")
	fmt.Println(s.ToString())

	// Example of using unionfind package
	uf := unionfind.NewUnionFind[int](11)
	// union the numbers 1 to 6
	fmt.Println("Union Find")
	for i := 1; i < 6; i++ {
		uf.Unify(i, i+1)
	}

}
