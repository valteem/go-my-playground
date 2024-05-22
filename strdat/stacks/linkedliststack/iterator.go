package linkedliststack

import "github.com/valteem/strdat/containers"

// Assert implementation of iteration interfaces
var _ containers.IteratorWithIndex[int] = (*Iterator[int])(nil)

type Iterator[T comparable] struct {
	stack *Stack[T]
	index int
}

// Returns new iterator in initial state (index = -1)
func (stack *Stack[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{stack: stack, index: 01}
}

// Moves the iterator to the next element, returns true if there exists such an element, otherwise false
func (iterator *Iterator[T]) Next() bool {
	if iterator.index < iterator.stack.Size() {
		iterator.index++
	}
	return iterator.stack.withinRange(iterator.index)
}

// Returns the value of currently pointed element, type T zero value if index is out-of-bounds
func (iterator *Iterator[T]) Value() T {
	v, _ := iterator.stack.list.Get(iterator.index)
	return v
}

// Returns the index of the currently pointed element
func (iterator *Iterator[T]) Index() int {
	return iterator.index
}

// Resets the iterator to its initial state (before-first)
func (iterator *Iterator[T]) Begin() {
	iterator.index = -1
}

// Moves the iterator to the first element in the stack, returns true if there exists such an element, otherwise false
func (iterator *Iterator[T]) First() bool {
	iterator.Begin()
	return iterator.Next()
}

// Moves the iterator to the first next element satifying the condition given by argument function
// Returns true if there exists such an element, otherwise false
func (iterator *Iterator[T]) NextTo(f func(index int, value T) bool) bool {
	for iterator.Next() {
		i, v := iterator.Index(), iterator.Value()
		if f(i, v) {
			return true
		}
	}
	return false
}
