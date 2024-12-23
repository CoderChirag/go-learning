package iter

import (
	"fmt"
	"iter"
	"slices"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	}else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// Finite Iterator
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val){
				return
			}
		}
	}
}

// Infinite Iterator
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a){
				return
			}
			a, b = b, a+b
		}
	}
}

func IterExamples() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)

	for v := range lst.All() {
		fmt.Println(v)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		if n > 15 {
			break
		}
		fmt.Println(n)
	}
}