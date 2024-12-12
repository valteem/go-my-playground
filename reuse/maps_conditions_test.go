package reuse_test

import (
	"reflect"
	"testing"

	"github.com/valteem/reuse"
)

func TestExtractKeyValues(t *testing.T) {

	tests := []struct {
		m      map[string][]string
		key    string
		output map[string]string
	}{
		{
			m: map[string][]string{
				"food[vegs]":   []string{"potatoes", "onions"},
				"food[fruits]": []string{"apples", "oranges"},
			},
			key:    "food",
			output: map[string]string{"vegs": "potatoes", "fruits": "apples"},
		},
	}

	for _, tc := range tests {

		output, _ := reuse.ExtractKeyValues(tc.m, tc.key)

		if !reflect.DeepEqual(output, tc.output) {
			t.Errorf("extracting values of key %q from map %v:\nget\n%v\nexpect\n%v", tc.key, tc.m, output, tc.output)
		}
	}
}
