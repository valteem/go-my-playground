// https://www.cogsci.ed.ac.uk/~richard/utf-8.cgi
package reuse_test

import (
	"testing"
)

func TestStringLen(t *testing.T) {

	tests := []struct {
		input  string
		length int
	}{
		{"abc", 3},
		{"ȀȠϾ", 6},
		{"ကឃℴ", 9},
		{"🐀🐃🐄", 12},
	}

	for _, tc := range tests {
		if len(tc.input) != tc.length {
			t.Errorf("len(%s): get %d, expect %d", tc.input, len(tc.input), tc.length)
		}
	}

}
