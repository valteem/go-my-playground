package doublylinkedlist

import (
	"encoding/json"

	"github.com/valteem/strdat/containers"
)

// Assert implementation of serialization interfaces
var _ containers.JSONSerializer = (*List[string])(nil)
var _ containers.JSONDeserializer = (*List[int])(nil)

// Returns JSON representation of the list
func (list *List[T]) ToJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}

// Populates the list with values in JSON format
func (list *List[T]) FromJSON(input []byte) error {
	var elements []T
	err := json.Unmarshal(input, &elements)
	if err == nil {
		list.Clear()
		list.Add(elements...)
	}
	return err
}

// Implements json.Marshaler{}
func (list *List[T]) MarshalJSON() ([]byte, error) {
	return list.ToJSON()
}

// Implements json.Unmarshaler{}
func (list *List[T]) UnmarshalJSON(input []byte) error {
	return list.FromJSON(input)
}
