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

func TestStringSlicing(t *testing.T) {

	tests := []struct {
		input    string
		prefixFn func(s string) string
		prefix   string
		suffixFn func(s string) string
		suffix   string
	}{
		{
			input: "create:alert:*",
			prefixFn: func(s string) string {
				return s[:len(s)-1]
			},
			prefix: "create:alert:",
			suffixFn: func(s string) string {
				return string(s[len(s)-1])
			},
			suffix: "*",
		},
	}

	for _, tc := range tests {
		if prefix, suffix := tc.prefixFn(tc.input), tc.suffixFn(tc.input); prefix != tc.prefix || suffix != tc.suffix {
			t.Errorf("%s prefix/suffix: get %s/%s, expect %s/%s", tc.input, prefix, suffix, tc.prefix, tc.suffix)
		}
	}

}
