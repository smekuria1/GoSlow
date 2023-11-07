

# GoSlow
[![Go Test](https://github.com/smekuria1/GoSlow/actions/workflows/test.yaml/badge.svg)](https://github.com/smekuria1/GoSlow/actions/workflows/test.yaml) [![Go Reference](https://pkg.go.dev/badge/github.com/smekuria1/GoSlow.svg)](https://pkg.go.dev/github.com/smekuria1/GoSlow) [![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)



GoSlow is combination of multiple packages that each represent popular algorithms and data structures. The goal of this project is to provide a simple and easy to use library for learning and teaching purposes.


## Acknowledgements

 - [William Fiset Algorithms in Java](https://github.com/williamfiset/algorithms)



## Binary Search Tree

 - [Package binarysearchtree](https://github.com/smekuria1/GoSlow/tree/main/binarysearchtree)
 #### Usage/Examples

```go
import "github.com/smekuria1/GoSlow/binarysearchtree"
bst := binarysearchtree.NewBST[int]()
```
## Dynamic Array
 - [Package darray](https://github.com/smekuria1/GoSlow/tree/main/darray)
#### Usage/Examples

```go
import "github.com/smekuria1/GoSlow/darray"
dynamicarray := darray.NewDynamicArray[int]()
```
## Doubly Linked List

- [Package doublyLinkedList](https://github.com/smekuria1/GoSlow/tree/main/doublyLinkedList)
#### Usage/Examples

```go
import "github.com/smekuria1/GoSlow/doublyLinkedList"
dll := doublyLinkedList.NewDoublyLinkedList[int]()
```
## HashTable(Separate Chaining)

- [Package hashtable](https://github.com/smekuria1/GoSlow/tree/main/hashtable)
#### Usage/Examples

```go
import "github.com/smekuria1/GoSlow/hashtable"
ht := hashtable.NewHashTable[int, int](10, 0.75)
```
## Binary Heap Priority Queue

- [Package priorityqueue](https://github.com/smekuria1/GoSlow/tree/main/priorityqueue)
#### Usage/Examples

```go
import "github.com/smekuria1/GoSlow/priorityqueue"
pq := priorityqueue.NewBinaryHeapPQ[int]()
```
## Queue

- [Package queue](https://github.com/smekuria1/GoSlow/tree/main/queue)
#### Usage/Examples

```go
import "github.com/smekuria1/GoSlow/queue"
q := queue.NewQueue[int]()
```
## Stack

- [Package stack](https://github.com/smekuria1/GoSlow/tree/main/stack)
#### Usage/Examples

```go
import "github.com/smekuria1/GoSlow/stack"
s := stack.NewStack[int]()
```
## Union Find

- [Package stack](https://github.com/smekuria1/GoSlow/tree/main/unionfind)
#### Usage/Examples

```go
import "github.com/smekuria1/GoSlow/unionfind"
uf := unionfind.NewUnionFind[int](11)
```
## Running Tests

To run tests for all packages, run the following command in root folder

```bash
    go test ./...
```

To run tests for specific packages, run the following command 

```bash
    cd <package name>
    go test 
```


## Support

For support, email solmek18@gmail.com or create and issue/pr.


## License

[MIT](https://choosealicense.com/licenses/mit/)


## FAQ

#### Why do this?

I love golang and wanted to excercise my golang muscles while learning about cool data structures and algos.



## Roadmap

- Additional Data Structures and Algos like AVL Trees, Fenwick Trees ....

- Build Either a CLI or small web app to visually interact with the implementations

- General improvements to some of the implementations and better error handling

