package main

import (
	"fmt"
	"math/rand"

	// darray "github.com/smekuria1/GoSlow/dArray"
	doublyLinkedList "github.com/smekuria1/GoSlow/doublyLinkedList"
	"github.com/smekuria1/GoSlow/priorityqueue"
	"github.com/smekuria1/GoSlow/stack"
)

func main() {
	//array := darray.NewDynamicArray[int](2)
	//var array []int
	i := 0
	//for i < 10 {
	//array = append(array, rand.Intn(1000-0)+0)
	//i += 1
	//	}
	// fmt.Println("Size: ", array.Size())
	// fmt.Println("Array: ", array.ToString())

	dl := doublyLinkedList.NewDoublyLinkedList[int]()

	for i < 10 {
		num := rand.Intn(1000-0) + 0
		dl.Add(num)
		//fmt.Printf("Added num: %v\n", num)
		i += 1

	}
	//fmt.Println("Size: ", dl.Size())
	dl.Add(200)
	// fmt.Println("LinkedList", dl.ToString())
	// fmt.Println(dl.Contains(11111))
	// fmt.Println(dl.RemoveAt(6))

	// dl.Reverse()
	// fmt.Println("LinkedList", dl.ToString())

	stack := stack.NewStack[int]()
	j := 0
	for j < 10 {
		num := rand.Intn(1000-0) + 0
		stack.Push(num)
		// fmt.Printf("Added num: %v\n", num)
		j += 1

	}
	// fmt.Println("Size: ", stack.Size())
	// fmt.Println("Stack", stack.ToString())
	// fmt.Println(stack.Contains(11111))
	// fmt.Println(stack.Peek())
	// fmt.Println("Stack", stack.ToString())

	pq := priorityqueue.NewBinaryHeapPQ[int]()
	k := 0
	for k < 10 {
		num := rand.Intn(1000-0) + 0
		pq.Add(num)
		fmt.Printf("Added num: %v\n", num)
		k += 1

	}
	pq.Add(400)

	fmt.Println(pq.ToString())
	fmt.Println(pq.IsMinHeap(0))
	fmt.Println(pq.ToString())
	fmt.Println(*pq.Peek())
	fmt.Println(*pq.Poll())
	fmt.Println(pq.ToString())
	fmt.Println(pq.IsMinHeap(0))
	fmt.Println(pq.Contains(400))
}
