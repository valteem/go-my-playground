package arraystack

import "github.com/valteem/strdat/containers"

// Assert implementation of IteratorWithIndex{}, ReverseIteratorWithIndex{}
var _ containers.IteratorWithIndex[int] = (*Iterator[int])(nil)
var _ containers.ReverseIteratorWithIndex[string] = (*Iterator[string])(nil)

type Iterator[T comparable] struct {
	stack *Stack[T]
	index int
}

// Returns an iterator whose values can be fetched by an index
func (stack *Stack[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{stack: stack, index: -1}
}

// Moves the iterator to the next element, returns true if there exists such an element, otherwise false
func (iterator *Iterator[T]) Next() bool {
	if iterator.index < iterator.stack.Size() {
		iterator.index++
	}
	return iterator.stack.withinRange(iterator.index)
}

// Moves the iterator to the previous element, returns true if there exists such an element, otherwise false
func (iterator *Iterator[T]) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	return iterator.stack.withinRange(iterator.index)
}

// Returns the value of the currently pointed element, or type T zero value if iterator index is out-of-bounds
func (iterator *Iterator[T]) Value() T {
	v, _ := iterator.stack.list.Get(iterator.stack.list.Size() - iterator.index - 1) // LIFO
	return v
}

// Returns the index of the currently pointed element
func (iterator *Iterator[T]) Index() int {
	return iterator.index
}

// Moves the iterator to its initial state (before-first)
func (iterator *Iterator[T]) Begin() {
	iterator.index = -1
}

// Moves the iterator past the last element
func (iterator *Iterator[T]) End() {
	iterator.index = iterator.stack.Size()
}

// Moves the iterator to the first (top) element in the stack, returns true if there exists such an element, otherwise false
func (iterator *Iterator[T]) First() bool {
	iterator.Begin()
	return iterator.Next()
}

// Moves the iterator to the last (bottom) element, returns tru if there ixists such an element, otherwise false
func (iterator *Iterator[T]) Last() bool {
	iterator.End()
	return iterator.Prev()
}

// Moves the iterator from the currently pointed element to first next element for which the argument function returns true,
// returns true if there exists such an element, otherwise false
func (iterator *Iterator[T]) NextTo(f func(index int, value T) bool) bool {
	for iterator.Next() {
		if i, v := iterator.Index(), iterator.Value(); f(i, v) {
			return true
		}
	}
	return false
}

// Moves the iterator from the currently pointed element to first previous element for which the argument function returns true,
// returns true if there exists such an element, otherwise false
func (iterator *Iterator[T]) PrevTo(f func(index int, value T) bool) bool {
	for iterator.Prev() {
		if i, v := iterator.Index(), iterator.Value(); f(i, v) {
			return true
		}
	}
	return false
}
