package singlylinkedlist

import (
	"github.com/valteem/strdat/containers"
)

// Assert singly linked list complies with EnumerableWithIndex{} interface
var _ containers.EnumerableWithIndex[string] = (*List[string])(nil)

// Calls the given function once on each element
func (list *List[T]) Each(f func(index int, value T)) {
	iterator := list.Iterator()
	for iterator.Next() {
		f(iterator.Index(), iterator.Value())
	}
}

// Returns a new list containing values returned by given function called once for each element
func (list *List[T]) Map(f func(index int, value T) T) *List[T] {
	newList := &List[T]{}
	iterator := list.Iterator()
	for iterator.Next() {
		newList.Add(f(iterator.Index(), iterator.Value()))
	}
	return newList
}

// Returns a new list containing elements for which the given function returns `true`
func (list *List[T]) Select(f func(index int, value T) bool) *List[T] {
	newList := &List[T]{}
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			newList.Add(iterator.Value())
		}
	}
	return newList
}

// Returns `true` if the given function returns `true` for any single element of the list
func (list *List[T]) Any(f func(index int, value T) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return true
		}
	}
	return false
}

// Returns `true` if the given function returns `true` for all elements in the list, otherwise `false`
func (list *List[T]) All(f func(index int, value T) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if !f(iterator.Index(), iterator.Value()) {
			return false
		}
	}
	return true
}

// Returns the first (index, value) pair for which the given function returns `true`, otherwise (-1, zero value)
func (list *List[T]) Find(f func(index int, value T) bool) (int, T) {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return iterator.Index(), iterator.Value()
		}
	}
	var t T
	return -1, t
}
