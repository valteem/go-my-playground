package reuse_test

import (
	"encoding/base64"
	"slices"
	"testing"

	"github.com/valteem/reuse"
)

func TestMainDecodeAndSlice(t *testing.T) {

	tests := []struct {
		name      string
		input     string
		delimiter string
		output    []string
	}{
		{
			name:      "string delimited by colons",
			input:     "value1:value2:value3",
			delimiter: ":",
			output:    []string{"value1", "value2", "value3"},
		},
		{
			name:      "string delimited by colons, delimiter set to comma",
			input:     "value1:value2:value3",
			delimiter: ",",
			output:    []string{"value1:value2:value3"},
		},
		{
			name:      "string delimited by colons, delimiter set to pattern",
			input:     "value1:value2:value3",
			delimiter: "value",
			output:    []string{"", "1:", "2:", "3"}, // leading zero string
		},
		{
			name:      "delimiter longer than decoded input",
			input:     "value1:value2:value3",
			delimiter: "value1:value2:value3:",
			output:    []string{},
		},
	}

	for _, tc := range tests {
		decodedInput := base64.StdEncoding.EncodeToString([]byte(tc.input))
		output, _ := reuse.DecodeAndSlice(decodedInput, tc.delimiter)
		if !slices.Equal(output, tc.output) {
			t.Errorf("%s: get %v, expect %v", tc.name, output, tc.output)
		}
	}

}
