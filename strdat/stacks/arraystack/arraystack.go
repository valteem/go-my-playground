package arraystack

import (
	"fmt"
	"strings"

	"github.com/valteem/strdat/lists/arraylist"
)

// Holds stack elements in an arraylist
type Stack[T comparable] struct {
	list *arraylist.List[T]
}

// Returns new empty stack
func New[T comparable]() *Stack[T] {
	return &Stack[T]{list: arraylist.New[T]()}
}

// Adds an element on top of the stack
func (stack *Stack[T]) Push(value T) {
	stack.list.Add(value)
}

// Removes top element from the stack, and returns it and true, or type T zero value and false if the stack is empty
func (stack *Stack[T]) Pop() (value T, ok bool) {
	value, ok = stack.list.Get(stack.list.Size() - 1)
	stack.list.Remove(stack.list.Size() - 1)
	return
}

// Returns top element from the stack and true, or type T zero value and false if the stack is empty
func (stack *Stack[T]) Peek() (value T, ok bool) {
	return stack.list.Get(stack.list.Size() - 1)
}

// Returns true if the stack is empty, false otherwise
func (stack *Stack[T]) Empty() bool {
	return stack.list.Empty()
}

// Returns number of elements in the stack
func (stack *Stack[T]) Size() int {
	return stack.list.Size()
}

// Removes all elements from the stack“
func (stack *Stack[T]) Clear() {
	stack.list.Clear()
}

// Returns all alements from the stack in LIFO order
func (stack *Stack[T]) Values() []T {
	size := stack.list.Size()
	values := make([]T, size)
	for i := 1; i <= size; i++ {
		values[size-i], _ = stack.list.Get(i - 1)
	}
	return values
}

// Returns string representation of the stack
func (stack *Stack[T]) String() string {
	output := "ArrayStack\n"
	values := []string{}
	for _, v := range stack.list.Values() {
		values = append(values, fmt.Sprintf("%v", v))
	}
	output += strings.Join(values, ", ")
	return output
}

// If the index is withing the bound of the underlying list
// Used in the implementations of Iterator / ReverseIterator
func (stack *Stack[T]) withinRange(index int) bool {
	return index >= 0 && index < stack.list.Size()
}
