package model

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestProduct(t *testing.T) {

	tests := []struct {
		input  string
		output *Product
	}{
		{
			input: `{"description":"some very good product", "features":[{"name":"color", "value":"blue"}, {"name":"weight", "value":42}]}`,
			output: &Product{
				Id:          0,
				Description: "some very good product",
				Features: []Feature{
					Feature{Name: "color", Value: "blue"},
					Feature{Name: "weight", Value: float64(42)}, // numbers always decoded as float64
				},
			},
		},
		{
			input: `{"description":"some other product"}`,
			output: &Product{
				Id:          0,
				Description: "some other product",
				Features:    nil,
			},
		},
	}

	for _, tc := range tests {
		output := &Product{}
		err := json.Unmarshal([]byte(tc.input), output)
		if err != nil {
			t.Errorf("failed to decode input: %v", err)
			continue
		}
		if !reflect.DeepEqual(output, tc.output) {
			t.Errorf("decoding %q input:\nget\n%v\nexpect\n%v\n", tc.input, output, tc.output)
		}
	}
}
