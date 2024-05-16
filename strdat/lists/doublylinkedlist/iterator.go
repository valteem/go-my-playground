package doublylinkedlist

// Holds the state of the iteration
type Iterator[T comparable] struct {
	list    *List[T]
	index   int
	element *element[T]
}

// Returns a new Iterator
func (list *List[T]) Iterator() Iterator[T] {
	return Iterator[T]{list: list, index: -1, element: nil}
}

// Moves the iterator to the next element, returns true if there exists such an element, false otherwise
func (iterator *Iterator[T]) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	if !iterator.list.withinRange(iterator.index) { // moved past last element
		iterator.element = nil
		return false
	}
	if iterator.index != 0 {
		iterator.element = iterator.element.next
	} else {
		iterator.element = iterator.list.first
	}
	return true
}

// Moves the iterator to the previous element, returns true if there exists such an element, otherwise false
func (iterator *Iterator[T]) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	if !iterator.list.withinRange(iterator.index) { // moved past first element
		iterator.element = nil
		return false
	}
	if iterator.index == iterator.list.size-1 { // moved from right (index = list.size) to the last element
		iterator.element = iterator.list.last
	} else {
		iterator.element = iterator.element.prev
	}
	return true
}

// Returns the value of the currently pointed element
func (iterator *Iterator[T]) Value() T {
	if iterator.index < 0 || iterator.index >= iterator.list.size {
		var t T
		return t
	}
	return iterator.element.value
}

// Returns the index of the currently pointed element
func (iterator *Iterator[T]) Index() int {
	return iterator.index
}

// Moves the iterator to before-first position
func (iterator *Iterator[T]) Begin() {
	iterator.index = -1
	iterator.element = nil
}

// Moves the iterator to after-last position
func (iterator *Iterator[T]) End() {
	iterator.index = iterator.list.size
	iterator.element = nil
}

// Moves the iterator to the first element, returns true if there exists such an element, otherwise false
func (iterator *Iterator[T]) First() bool {
	iterator.Begin()
	return iterator.Next()
}

// Moves the iterator to the last element, returns true if there exists such an element, otherwise false
func (iterator *Iterator[T]) Last() bool {
	iterator.End()
	return iterator.Prev()
}

// Moves the iterator from its current position to the next element that satisfies the condition defined by argument function,
// returns true if there exists such an element, moves past last element abd returns false otherwise
func (iterator *Iterator[T]) NextTo(f func(index int, value T) bool) bool {
	for iterator.Next() {
		i, v := iterator.Index(), iterator.Value()
		if f(i, v) {
			return true
		}
	}
	return false
}

// Moves the iterator from its current position to the previous element that satisfies condition defined by argument function,
// returns true if there exists such an element, moves past first element and returns false otherwise
func (iterator *Iterator[T]) PrevTo(f func(index int, value T) bool) bool {
	for iterator.Prev() {
		i, v := iterator.Index(), iterator.Value()
		if f(i, v) {
			return true
		}
	}
	return false
}
