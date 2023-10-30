// Package UnionFind implements the union find data structure.
// uses my darray package for the union find
package unionfind

import (
	"fmt"
	"log"

	"github.com/smekuria1/GoSlow/darray"
)

type UnionFind[T comparable] struct {
	size          int
	sz            *darray.DynamicArray[int]
	id            *darray.DynamicArray[int]
	numComponents int
}

func NewUnionFind[T comparable](size int) *UnionFind[T] {
	if size <= 0 {
		log.Println("size must be > 0 returning empty UnionFind")
		return nil
	}
	sz := darray.NewDynamicArray[int](size)
	id := darray.NewDynamicArray[int](size)
	for i := 0; i < size; i++ {
		id.Add(i)
		sz.Add(1)
	}

	return &UnionFind[T]{
		size:          size,
		sz:            sz,
		id:            id,
		numComponents: size,
	}
}

// Find which component/set 'p' belongs to, takes amortized constant time.
func (uf *UnionFind[T]) Find(p int) int {
	root := p

	for root != uf.id.Get(root) {
		root = uf.id.Get(root)
	}

	// path compression
	for p != root {
		next := uf.id.Get(p)
		uf.id.Set(p, root)
		p = next
	}
	fmt.Println(uf.id.ToString())
	return root
}

// Connected returns whether p and q are in the same component/set
func (uf *UnionFind[T]) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// ComponentSize returns the size of the component/set p belongs to
func (uf *UnionFind[T]) ComponentSize(p int) int {
	return uf.sz.Get(uf.Find(p))
}

// Components returns the number of remaining components/sets
func (uf *UnionFind[T]) Components() int {
	return uf.numComponents
}

// Size returns the number of elements in this UnionFind/Disjoint set
func (uf *UnionFind[T]) Size() int {
	return uf.size
}

// Unify unifies the components containing elements p and q
// not the best approach since i used my own darray impl
func (uf *UnionFind[T]) Unify(p, q int) {
	//check if elements are already connected
	if uf.Connected(p, q) {
		return
	}

	root1 := uf.Find(p)
	root2 := uf.Find(q)

	//merge smaller component/sets into larger one
	//very convoluted because i used my darray implementation
	if uf.sz.Get(root1) < uf.sz.Get(root2) {
		uf.sz.Set(root2, uf.sz.Get(root2)+uf.sz.Get(root1))
		uf.id.Set(root2, root1)
		uf.sz.Set(root1, 0)
	} else {
		uf.sz.Set(root1, uf.sz.Get(root1)+uf.sz.Get(root2))
		uf.id.Set(root2, root1)
		uf.sz.Set(root2, 0)
	}

	uf.numComponents--
	// print all the roots and what they point to
	fmt.Println(uf.id.ToString())
	fmt.Println(uf.sz.ToString())
	fmt.Println(uf.numComponents)

}
