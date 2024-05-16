package doublylinkedlist

import "github.com/valteem/strdat/containers"

// Assert Enumerable{} implementation
var _ containers.EnumerableWithIndex[int] = (*List[int])(nil)

// Calls the function once for each element, taking index and value as arguments
func (list *List[T]) Each(f func(index int, value T)) {
	iterator := list.Iterator()
	for iterator.Next() {
		f(iterator.Index(), iterator.Value())
	}
}

// Invokes the argument function once for each element, returns new list containing the values returned by the ergument function
func (list *List[T]) Map(f func(index int, value T) T) *List[T] {
	new := &List[T]{}
	iterator := list.Iterator()
	for iterator.Next() {
		new.Add(f(iterator.Index(), iterator.Value()))
	}
	return new
}

// Returns a new list containing all elements for which the argument function returns true
func (list *List[T]) Select(f func(index int, value T) bool) *List[T] {
	new := &List[T]{}
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			new.Add(iterator.Value())
		}
	}
	return new
}

// Returns true if the argument function returns true for any element of the list, otherwise false
func (list *List[T]) Any(f func(index int, value T) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return true
		}
	}
	return false
}

// Returns true if the argument function returns true for all elements of the list, otherwise false
func (list *List[T]) All(f func(index int, value T) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if !f(iterator.Index(), iterator.Value()) {
			return false
		}
	}
	return true
}

// Returns index and value of the first element for which the argument function returns true,
// and (-1, zero value of type T) if no such element exists
func (list *List[T]) Find(f func(index int, value T) bool) (index int, value T) {
	iterator := list.Iterator()
	for iterator.Next() {
		if i, v := iterator.Index(), iterator.Value(); f(i, v) {
			return i, v
		}
	}
	var t T
	return -1, t
}
