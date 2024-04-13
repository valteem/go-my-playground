package singlylinkedlist

import (
	"encoding/json"

	"github.com/valteem/strdat/containers"
)

// Assert implementation of basic serialization/deserialization interfaces
var _ containers.JSONSerializer = (*List[int])(nil)
var _ containers.JSONDeserializer = (*List[string])(nil)

func (list *List[T]) ToJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}

// Overwrites list contents with values contained in `input“
func (list *List[T]) FromJSON(input []byte) error {
	var elements []T
	err := json.Unmarshal(input, &elements)
	if err == nil {
		list.Clear()
		list.Add(elements...)
	}
	return err
}

func (list *List[T]) MarshalJSON() ([]byte, error) {
	return list.ToJSON()
}

func (list *List[T]) UnmarshalJSON(input []byte) error {
	return list.FromJSON(input)
}
