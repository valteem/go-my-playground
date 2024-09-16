package reuse_test

import (
	"testing"

	"github.com/valteem/reuse"
)

/*
    var (
	    nilSlice []int
    )
*/

func TestNilCheck(t *testing.T) {

	tests := []struct {
		name   string
		input  any
		output bool
	}{
		{
			name: "anonymous struct",
			input: struct {
				key   int
				value string
			}{},
			output: false,
		},
		{
			name:   "slice (in place)",
			input:  []int{},
			output: false,
		},
		// doesn't work on nil slices
		/* 		{
			name:   "slice (var)",
			input:  nilSlice,
			output: true,
		}, */
	}

	for _, tc := range tests {
		if output := reuse.NilCheck(tc.input); output != tc.output {
			t.Errorf("%s: get %t, expect %t", tc.name, output, tc.output)
		}
	}

}
