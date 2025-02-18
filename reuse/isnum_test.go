package reuse_test

import (
	"slices"
	"testing"

	"github.com/valteem/reuse"
)

func TestIsNum(t *testing.T) {

	tests := []struct {
		input  string
		output []bool
	}{
		{
			input:  "0123456789",
			output: []bool{true, true, true, true, true, true, true, true, true, true},
		},
		{
			input:  "1a5A&%",
			output: []bool{true, false, true, false, false, false},
		},
	}

	for _, tc := range tests {

		output := []bool{}

		for _, v := range tc.input {
			output = append(output, reuse.IsNum(v))
		}

		if !slices.Equal(output, tc.output) {
			t.Errorf("calculate IsNum() for every digit in %q:\nget\n%v\nexpect\n%v\n", tc.input, output, tc.output)
		}
	}

}
