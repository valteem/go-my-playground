package arraylist

import (
	"github.com/valteem/strdat/containers"
)

// Assert implementation of interfaces from containers/iterator.go
var _ containers.IteratorWithIndex[string] = (*Iterator[string])(nil) // dereferencing pointer (interface methods implemented with pointer receiver) + type conversion (https://stackoverflow.com/a/69470411)
var _ containers.ReverseIteratorWithIndex[string] = (*Iterator[string])(nil)

// Holds the state of iterator
type Iterator[T comparable] struct {
	list  *List[T]
	index int
}

// Returns stateful iterator, its values can be fetched by index
// Arbitrary initial value (-1) for index
func (list *List[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{list: list, index: -1}
}

// Moves the iterator to the next element. Stops at list.size
func (iterator *Iterator[T]) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	return iterator.list.withinRange(iterator.index)
}

// Moves the iterator to the previous element. Stops at -1
func (iterator *Iterator[T]) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	return iterator.list.withinRange(iterator.index)
}

func (iterator *Iterator[T]) Index() int {
	return iterator.index
}

func (iterator *Iterator[T]) Value() T {
	return iterator.list.elements[iterator.index]
}

func (iterator *Iterator[T]) Begin() {
	iterator.index = -1
}

func (iterator *Iterator[T]) End() {
	iterator.index = iterator.list.size
}

func (iterator *Iterator[T]) First() bool {
	iterator.Begin()
	return iterator.Next()
}

func (iterator *Iterator[T]) Last() bool {
	iterator.End()
	return iterator.Prev()
}

// Returns next to current position satisfying given condition
func (iterator *Iterator[T]) NextTo(f func(index int, value T) bool) bool {
	for iterator.Next() {
		i, v := iterator.Index(), iterator.Value()
		if f(i, v) {
			return true
		}
	}
	return false
}

// Returns previous to current position satisfying given condition
func (iterator *Iterator[T]) PrevTo(f func(index int, value T) bool) bool {
	for iterator.Prev() {
		i, v := iterator.Index(), iterator.Value()
		if f(i, v) {
			return true
		}
	}
	return false
}
