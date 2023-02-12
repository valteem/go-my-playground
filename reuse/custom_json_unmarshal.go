// https://mariadesouza.com/2017/09/07/custom-unmarshal-json-in-golang/
package reuse

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title string `json:"title"`
	Size string `json:"size"`
	Pages int16 `json:"pages"`
}

type Order struct {
	Ref string `json:"ref"`
	Books []Book `json:"books"`
}

type ShortBook struct {
	Title string `json:"title"`
	Size string `json:"size"`
	Pages int16 `json:"pages"`
}

type ShortOrder struct {
	Ref string `json:"ref"`
	Books []ShortBook `json:"books"`
}

func OrderUnmatshal( b []byte) *Order {
	u := Order{}
	err := json.Unmarshal(b, &u)
	if err != nil {
		fmt.Println(err)
		return &Order{}
	}
	return &u
}

func (s *ShortBook) UnmarshalJSON(b []byte) error {
	var v []interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	s.Title, _ = v[0].(string)
	s.Size, _ = v[1].(string)
	s.Pages = int16(v[2].(float64)) // v[2].(int16) return 0

	return nil
}

func ShortOrderUnmatshal( b []byte) *ShortOrder {
	u := ShortOrder{}
	err := json.Unmarshal(b, &u)
	if err != nil {
		fmt.Println(err)
		return &ShortOrder{}
	}
	return &u
}