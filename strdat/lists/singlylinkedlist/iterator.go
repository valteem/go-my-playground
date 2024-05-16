package singlylinkedlist

import (
	"github.com/valteem/strdat/containers"
)

// Assert iterator implements interfaces defined in containers/iterator
var _ containers.IteratorWithIndex[int] = (*Iterator[int])(nil)

// Not implemented
// var _ containers.ReverseIteratorWithIndex[int] = (*Iterator[int])(nil)

type Iterator[T comparable] struct {
	list    *List[T]
	index   int
	element *element[T]
}

func (list *List[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{list: list, index: -1, element: nil} // index: -1 - one-before-first element state
}

// Moves the iterator to the next element, both index and pointer to element change
// Returns `true` if there exists such a non-nil element, otherwise `false`
// Calling Next() for the first time moves the iterator to the first element
func (iterator *Iterator[T]) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	if !iterator.list.withinRange(iterator.index) { // moving beyond last element, index set to `list.size``
		iterator.element = nil
		return false
	}
	if iterator.index == 0 {
		iterator.element = iterator.list.first
	} else {
		iterator.element = iterator.element.next
	}
	return true
}

// Returns value of the current element
func (iterator *Iterator[T]) Value() T {
	if iterator.index < 0 || iterator.index >= iterator.list.size {
		var t T
		return t
	}
	return iterator.element.value
}

// Returns index of the current element
func (iterator Iterator[T]) Index() int {
	return iterator.index
}

// Resets the iterator to its initial state (after list.Iterator())
func (iterator *Iterator[T]) Begin() {
	iterator.index = -1
	iterator.element = nil
}

// Moves the iterator to the first element in list
// Returns `true` if there exists first element, otherwise `false`
func (iterator *Iterator[T]) First() bool {
	iterator.Begin()
	return iterator.Next()
}

// Moves iterator from current position to first next element satisfying the condition given by argument function
// Returns `true` if such an element exists, otherwise `false`
func (iterator *Iterator[T]) NextTo(f func(index int, value T) bool) bool {
	for iterator.Next() {
		i, v := iterator.Index(), iterator.Value()
		if f(i, v) {
			return true
		}
	}
	return false
}
