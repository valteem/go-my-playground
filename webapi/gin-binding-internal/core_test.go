package main

import (
	"testing"
)

func TestHeadTail(t *testing.T) {

	tests := []struct {
		input string
		sep   string
		head  string
		tail  string
	}{
		{"more sweet apples", "app", "more sweet ", "les"},
		{"know your customer", "x", "know your customer", ""},
	}

	for _, tc := range tests {
		if head, tail := HeadTail(tc.input, tc.sep); head != tc.head || tail != tc.tail {
			t.Errorf("head/tail for %s:%s: get  %s/%s, expect %s/%s", tc.input, tc.sep, head, tail, tc.head, tc.tail)
		}
	}
}
