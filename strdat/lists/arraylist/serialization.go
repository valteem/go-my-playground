package arraylist

import (
	"encoding/json"

	"github.com/valteem/strdat/containers"
)

var _ containers.JSONSerializer = (*List[int])(nil)
var _ containers.JSONDeserializer = (*List[int])(nil)

func (list *List[T]) ToJSON() ([]byte, error) {
	return json.Marshal(list.elements[:list.size])
}

func (list *List[T]) FromJSON(input []byte) error {
	err := json.Unmarshal(input, &list.elements)
	if err == nil {
		list.size = len(list.elements)
	}
	return err
}

func (list *List[T]) UnmarshalJSON(bytes []byte) error {
	return list.FromJSON(bytes)
}

func (list *List[T]) MarshalJSON() ([]byte, error) {
	return list.ToJSON()
}
