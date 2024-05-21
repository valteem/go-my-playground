package arraystack

import "github.com/valteem/strdat/containers"

// Assert implementation of serialization intefaces
var _ containers.JSONSerializer = (*Stack[int])(nil)
var _ containers.JSONDeserializer = (*Stack[string])(nil)

// Returns JSON representation of the underlying list
func (stack *Stack[T]) ToJSON() ([]byte, error) {
	return stack.list.ToJSON() // FIFO
}

// Populates the stack with values from JSON representation for the underlying list
func (stack *Stack[T]) FromJSON(input []byte) error {
	return stack.list.FromJSON(input)
}

// Implements json.Marshaler
func (stack *Stack[T]) MarshalJSON() ([]byte, error) {
	return stack.ToJSON()
}

// Implements json.Unmarshaler
func (stack *Stack[T]) UnmarshalJSON(input []byte) error {
	return stack.FromJSON(input)
}
