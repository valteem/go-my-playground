package singlylinkedlist

import (
	"fmt"
	"slices"
	"strings"

	"github.com/valteem/strdat/lists"
	"github.com/valteem/strdat/utils"
)

// Assert correct implementation of List{} interface
var _ lists.List[int] = (*List[int])(nil)

type element[T comparable] struct {
	value T
	next  *element[T]
}

type List[T comparable] struct {
	first *element[T]
	last  *element[T]
	size  int
}

func New[T comparable](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List[T]) Add(values ...T) {
	for _, v := range values {
		new := &element[T]{value: v}
		if list.size == 0 {
			list.first = new
			list.last = new
		} else {
			list.last.next = new
			list.last = new
		}
		list.size++
	}
}

func (list *List[T]) Append(values ...T) {
	list.Add(values...)
}

func (list *List[T]) Prepend(values ...T) {
	// prepend one by one in reversed order to keep input order
	for i := len(values) - 1; i >= 0; i-- {
		new := &element[T]{value: values[i], next: list.first}
		list.first = new
		if list.size == 0 {
			list.last = new
		}
		list.size++
	}
}

func (list *List[T]) Get(index int) (T, bool) {
	if !list.withinRange(index) {
		var t T
		return t, false
	}
	element := list.first
	for i := 0; i != index; i, element = i+1, element.next {
	}
	return element.value, true
}

func (list *List[T]) Remove(index int) {
	if !list.withinRange(index) {
		return
	}
	if list.size == 1 {
		list.Clear()
	}
	var beforeElement *element[T]
	element := list.first
	for i := 0; i != index; i, element = i+1, element.next {
		beforeElement = element
	}
	if element == list.first {
		list.first = element.next
	}
	if element == list.last {
		list.last = beforeElement
	}
	if beforeElement != nil { // beforeElement == nil if index == 0 (removing first (head) element)
		beforeElement.next = element.next
	}
	element = nil
	list.size--
}

func (list *List[T]) Contains(values ...T) bool {
	if len(values) == 0 {
		return true // any set always contains empty set
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
		if !found {
			return false // quit if any value from `values` is not found
		}
	}
	return true
}

func (list *List[T]) Values() []T {
	values := make([]T, list.size)
	for i, e := 0, list.first; e != nil; i, e = i+1, e.next {
		values[i] = e.value
	}
	return values
}

// Returns index if element (value) is found, otherwise -1
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

// Returns `true` if list does not contain any elements
func (list *List[T]) Empty() bool {
	return list.size == 0
}

// Accessor method for list.size
func (list *List[T]) Size() int {
	return list.size
}

func (list *List[T]) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

// In-place sort
func (list *List[T]) Sort(comparator utils.Comparator[T]) {
	if list.size < 2 {
		return
	}
	values := list.Values()
	slices.SortFunc(values, comparator)
	list.Clear()
	list.Add(values...)
}

func (list *List[T]) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) && i != j {
		var e1, e2 *element[T]
		for k, current := 0, list.first; e1 == nil || e2 == nil; k, current = k+1, current.next {
			switch k {
			case i:
				e1 = current
			case j:
				e2 = current
			}
		}
		e1.value, e2.value = e2.value, e1.value
	}
}

// Shifts right elements starting from given position (including this position)
// Does nothing if position is outside list body (negative or greater than list.size)
// Appends if position if equal to list.size
func (list *List[T]) Insert(index int, values ...T) {
	if len(values) == 0 {
		return // nothing to insert
	}
	if !list.withinRange(index) {
		if index == list.size { // append
			list.Add(values...)
		}
		return
	}
	list.size += len(values)
	var beforeElement *element[T]
	foundElement := list.first
	for i := 0; i != index; i, foundElement = i+1, foundElement.next {
		beforeElement = foundElement
	}
	if foundElement == list.first { // looks like this case is equivalent to index == 0
		oldNextElement := list.first
		for j, v := range values {
			newElement := &element[T]{value: v}
			if j == 0 {
				list.first = newElement
			} else {
				// no chance (?) for beforeElement == nil since first value from `values` has already been added
				beforeElement.next = newElement
			}
			beforeElement = newElement
		}
		beforeElement.next = oldNextElement
	} else {
		oldNextElement := beforeElement.next
		for _, v := range values {
			newElement := &element[T]{value: v}
			beforeElement.next = newElement
			beforeElement = newElement
		}
		beforeElement.next = oldNextElement
	}
}

// Set element value at given position (index)
// Do nothing if index is negative of greater than list.size
// Append if index is equal to list.size
func (list *List[T]) Set(index int, value T) {
	if !list.withinRange(index) {
		if index == list.size { // append
			list.Add(value)
		}
		return
	}
	foundElement := list.first
	for i := 0; i != index; {
		i, foundElement = i+1, foundElement.next
	}
	foundElement.value = value
}

// Returns string representation
func (list *List[T]) String() string {
	str := "SinglyLinkedList\n"
	values := []string{}
	for e := list.first; e != nil; e = e.next {
		values = append(values, fmt.Sprintf("%v", e.value))
	}
	str += strings.Join(values, ", ")
	str = strings.TrimSuffix(str, " ")
	return str
}

func (list *List[T]) withinRange(index int) bool {
	return index >= 0 && index < list.size
}
