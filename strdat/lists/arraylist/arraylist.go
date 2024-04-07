package arraylist

import (
	"fmt"
	"slices"
	"strings"

	"github.com/valteem/strdat/utils"
)

type List[T comparable] struct {
	elements []T
	size     int
}

const (
	growthFactor = float32(2.0)
	shrinkFactor = float32(0.25) // shrink when size reaches 1/4 of capacity
)

func (list *List[T]) resize(cap int) {
	newElements := make([]T, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

func (list *List[T]) growBy(n int) {
	currentCapacity := cap(list.elements)
	if (list.size + n) >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

func (list *List[T]) Add(values ...T) {
	list.growBy(len(values))
	for _, v := range values {
		list.elements[list.size] = v
		list.size++
	}
}

func New[T comparable](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List[T]) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

// First returned value is the element at the given index
// (provided the index is within bounds of underlying `array`, otherwise `initial` value for type T)
// Second returned  value is `true` if the index is within bounds of underlying `array`, otherwise `false`
func (list *List[T]) Get(index int) (T, bool) {
	if !list.withinRange(index) {
		var t T
		return t, false
	}
	return list.elements[index], true
}

// Shrinks the list if size is shrinkFactor of current capacity
func (list *List[T]) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}

// Removes the element at the given index
func (list *List[T]) Remove(index int) {
	if !list.withinRange(index) {
		return // (?) silently returns without throwing error or returning `false` as second return value
	}
	clear(list.elements[index : index+1]) // what for? this just `nullifies` list.elements[index]
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	list.shrink()
}

// Returns `true` if ALL elements from `values` are present in the list
// Also returns `true` if no arguments are passed
// Worst case performance  M x N, M = len(values), N = list.size
func (list *List[T]) Contains(values ...T) bool {
	for _, v := range values {
		found := false
		for i := 0; i < list.size; i++ {
			if list.elements[i] == v {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Returns all elements in the list
func (list *List[T]) Values() []T {
	output := make([]T, list.size)
	copy(output, list.elements[:list.size])
	return output
}

// Returns index of the given element, -1 if the element is not in the list
// (?) Does not seem to be a part of any interface, and is exported at the same time
func (list *List[T]) IndexOf(value T) int {
	if list.size == 0 {
		return -1
	}
	for i, elt := range list.elements {
		if elt == value {
			return i
		}
	}
	return -1
}

func (list *List[T]) Empty() bool {
	return list.size == 0
}

func (list *List[T]) Size() int {
	return list.size
}

func (list *List[T]) Clear() {
	list.size = 0
	list.elements = []T{}
}

func (list *List[T]) Sort(comparator utils.Comparator[T]) {
	if len(list.elements) < 2 {
		return
	}
	slices.SortFunc(list.elements[:list.size], comparator) // avoid sorting pre-allocated `sleeping` elements
}

func (list *List[T]) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

// Inserts value before the given position (index), valid positions are 0 (insert at the beginning), 1, 2, ..., list.size (append)
// Does nothing if position (index) is negative or greater than list size
func (list *List[T]) Insert(index int, values ...T) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(values...)
		}
		return
	}
	l := len(values)
	list.growBy(l)
	list.size += l
	copy(list.elements[index+l:], list.elements[index:list.size-1]) // (list.size - l) - end of list before growBy(l)
	copy(list.elements[index:], values)
}

// Sets the value of an element at the given position (index), valid positions are 0, 1, ..., list.size (append)
// Does nothing if position is negative of greater than list size
func (list *List[T]) Set(index int, value T) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(value)
		}
		return
	}
	list.elements[index] = value
}

// String representation
func (list *List[T]) String() string {
	str := "ArrayList\n"
	values := make([]string, 0, list.size)
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	str = strings.TrimSuffix(str, " ")
	return str
}
