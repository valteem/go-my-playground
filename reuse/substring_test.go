package reuse_test

import (
	"testing"

	"github.com/valteem/reuse"
)

func TestFindLongestSubstring(t *testing.T) {

	tests := []struct {
		input  string
		length int
	}{
		{"aabbcc", 2},
		{"aabcabbc", 3},
		{"onionsforsugar", 7},
	}

	for _, tc := range tests {
		if l := reuse.FindLongestSubstring(tc.input); l != tc.length {
			t.Errorf("%q: get %d, expect %d", tc.input, l, tc.length)
		}
	}

}
