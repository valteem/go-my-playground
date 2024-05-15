package doublylinkedlist

import (
	"fmt"
	"slices"
	"strings"

	"github.com/valteem/strdat/utils"
)

type element[T comparable] struct {
	value T
	prev  *element[T]
	next  *element[T]
}

type List[T comparable] struct {
	first *element[T]
	last  *element[T]
	size  int
}

// Instantiate a list and add values (if provided)
func New[T comparable](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Append values (one or more) at the end of the list (same as Append())
func (list *List[T]) Add(values ...T) {
	for _, value := range values {
		newElement := &element[T]{value: value, prev: list.last} // newElement.next = nil
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

// Append values (one or more) at the end of the list (same as Add())
func (list *List[T]) Append(values ...T) {
	list.Add(values...)
}

// Prepend values (one or more)
func (list *List[T]) Prepend(values ...T) {
	// adding in reverse order, to keep passed order
	for i := len(values) - 1; i >= 0; i-- { // (i >= 0) ensures no further actions if len(values) == 0
		newElement := &element[T]{value: values[i], next: list.first} // newElement.prev = nil
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.first.prev = newElement
			list.first = newElement
		}
		list.size++
	}
}

// Returns the element at index provided, and true if index is within list bounds, otherwise false
func (list *List[T]) Get(index int) (T, bool) {
	if !list.withinRange(index) {
		var t T
		return t, false
	}
	if list.size-index < index { // traverse from the end of the list
		element := list.last
		for i := list.size - 1; i != index; i, element = i-1, element.prev {
		}
		return element.value, true
	} else { // traverse from the beginning of the list
		element := list.first
		for i := 0; i != index; i, element = i+1, element.next {
		}
		return element.value, true
	}
}

// Removes the elemend at the index provided
func (list *List[T]) Remove(index int) {
	if !list.withinRange(index) {
		return
	}
	if list.size == 1 {
		list.Clear()
		return
	}
	var element *element[T]
	if list.size-index < index { // traverse from the end of the list
		element = list.last
		for i := list.size - 1; i != index; i, element = i-1, element.prev {
		}
	} else { // traverse from the beginning of the list
		element = list.first
		for i := 0; i != index; i, element = i+1, element.next {
		}
	}
	if element == list.first {
		list.first = element.next
	} else {
		element.prev.next = element.next // also ensures new list.last.next = nil when index = len(list) - 1
	}
	if element == list.last {
		list.last = element.prev
	} else {
		element.next.prev = element.prev // also ensures new list.first.prev = nil when index = 0
	}
	list.size--
}

// Removes all elements from the list
func (list *List[T]) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

// Returns true if all values are in the list, otherwise false
func (list *List[T]) Contains(values ...T) bool {
	if len(values) == 0 { // empty set is always a part of any other set
		return true
	}
	if list.size == 0 {
		return false
	}
	for _, v := range values {
		found := false
		for e := list.first; e != nil; e = e.next {
			if e.value == v {
				found = true
				break
			}
		}
		if !found { // at least one value is not in the list
			return false
		}
	}
	return true
}

// Returns values of all elements in the list as slice
func (list *List[T]) Values() []T {
	values := make([]T, list.size)
	for i, e := 0, list.first; e != nil; i, e = i+1, e.next {
		values[i] = e.value
	}
	return values
}

// Returns index of the  provided element, (-1) if not found
func (list *List[T]) IndexOf(value T) int {
	if list.size == 0 {
		return -1
	}
	for i, v := range list.Values() {
		if v == value {
			return i
		}
	}
	return -1
}

// True of the list does not contain any elements, false otherwise
func (list *List[T]) Empty() bool {
	return list.size == 0
}

// Number of elements in the list
func (list *List[T]) Size() int {
	return list.size
}

// Sorts list values in-place
func (list *List[T]) Sort(comp utils.Comparator[T]) {
	if list.size < 2 {
		return
	}
	values := list.Values()
	slices.SortFunc(values, comp)
	list.Clear()
	list.Add(values...)
}

// Swaps two elelents at the provided positions
func (list *List[T]) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		var e1, e2 *element[T]
		for k, e := 0, list.first; e1 == nil || e2 == nil; k, e = k+1, e.next {
			switch k {
			case i:
				e1 = e
			case j:
				e2 = e
			}
		}
		e1.value, e2.value = e2.value, e1.value
	}
}

// Inserts provided values at specified position, moving next elements (if any) to the right. Inserting at list.size appends values.
func (list *List[T]) Insert(index int, values ...T) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(values...)
		}
		return
	}
	var before, found *element[T]
	if list.size-index < index { // traverse from the end of the list
		found = list.last
		before = found.prev
		for i := list.size - 1; i != index; i, found = i-1, found.prev {
			before = before.prev
		}
	} else { // traverse from the beginning of the list
		found = list.first
		for i := 0; i != index; i, found = i+1, found.next {
			before = found
		}
	}
	if found == list.first {
		next := list.first // will be next to last inserted element
		for i, v := range values {
			new := &element[T]{value: v}
			if i == 0 {
				list.first = new
			} else {
				new.prev = before
				before.next = new
			}
			before = new
		}
		next.prev = before
		before.next = next
	} else {
		next := before.next // will be next to last inserted element
		for _, v := range values {
			new := &element[T]{value: v}
			new.prev = before
			before.next = new
			before = new
		}
		next.prev = before
		before.next = next
	}
	list.size++
}

// Sets value at specified position. Does nothing if index is out-of-bounds. Set at list.size appends an element.
func (list *List[T]) Set(index int, value T) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(value)
		}
		return
	}
	var found *element[T]        // found element belongs to scope that includes both if/else scopes
	if list.size-index < index { // traverse from the end of the list
		found = list.last
		for i := list.size - 1; i != index; i, found = i-1, found.prev {
		}
	} else { // traverse from the beginning of the list
		found = list.first
		for i := 0; i != index; i, found = i+1, found.next {
		}
	}
	found.value = value
}

// String representation of the list
func (list *List[T]) String() string {
	output := "DoublyLinkedList\n"
	strValues := []string{}
	for e := list.first; e != nil; e = e.next {
		strValues = append(strValues, fmt.Sprintf("%v", e.value))
	}
	output += strings.Join(strValues, ", ")
	output = strings.TrimSuffix(output, ", ")
	return output
}

// If the index provided is within bounds of the list
func (list *List[T]) withinRange(index int) bool {
	return index >= 0 && index < list.size
}
