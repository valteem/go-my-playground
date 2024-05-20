package linkedliststack

import (
	"fmt"
	"strings"

	"github.com/valteem/strdat/lists/singlylinkedlist"
	"github.com/valteem/strdat/stacks"
)

// Assert Stack implementation
var _ stacks.Stack[string] = (*Stack[string])(nil)

// Holds its elements in singly linked list
type Stack[T comparable] struct {
	list *singlylinkedlist.List[T]
}

// Returns new empty stack
func New[T comparable]() *Stack[T] {
	return &Stack[T]{list: singlylinkedlist.New[T]()}
}

// Adds a value on top of the stack
func (stack *Stack[T]) Push(value T) {
	stack.list.Prepend(value)
}

// Removes top element from the stack, returns it and true if the stack is not empty, or type T zero value and false otherwise
func (stack *Stack[T]) Pop() (T, bool) {
	value, ok := stack.list.Get(0)
	stack.list.Remove(0)
	return value, ok
}

// Returns top elelent and true if the stack is not empty, or type T zero value and false otherwise
func (stack *Stack[T]) Peek() (T, bool) {
	return stack.list.Get(0)
}

// Returns true on empty stack, false otherwise
func (stack *Stack[T]) Empty() bool {
	return stack.list.Empty()
}

// Returns number of elements in the stack
func (stack *Stack[T]) Size() int {
	return stack.list.Size()
}

// Removes all elements from the stack
func (stack *Stack[T]) Clear() {
	stack.list.Clear()
}

// Returns all elements in the stack
func (stack *Stack[T]) Values() []T {
	return stack.list.Values()
}

// Returns string representation of the stack
func (stack *Stack[T]) String() string {
	output := "LinkedListStack\n"
	values := []string{}
	for _, v := range stack.list.Values() {
		values = append(values, fmt.Sprintf("%v", v))
	}
	output += strings.Join(values, ", ")
	return output
}

// Returns true if index is withing the bounds of the underlying list, false otherwise
func (stack *Stack[T]) withinRange(index int) bool {
	return index >= 0 && index < stack.list.Size()
}
