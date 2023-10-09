package main

import (
	"fmt"
	"math/rand"

	// darray "example.com/goslow/darray"
	doublyLinkedList "github.com/smekuria1/GoSlow/doublyLinkedList"
)

func main() {
	// array := darray.NewDynamicArray(10)
	i := 0
	// for i < 5 {
	// 	array.Add(rand.Intn(1000-0) + 0)
	// 	i += 1
	// }
	dl := doublyLinkedList.NewDoublyLinkedList[int]()
	for i < 10 {
		num := rand.Intn(1000-0) + 0
		dl.Add(num)
		fmt.Printf("Added num: %v\n", num)
		i += 1

	}
	dl.AddLast(200)
	fmt.Println("Size: ", dl.Size())
	fmt.Println("LinkedList", dl.ToString())
	fmt.Println(dl.RemoveVal(9901))
	fmt.Println("LinkedList", dl.ToString())

}
