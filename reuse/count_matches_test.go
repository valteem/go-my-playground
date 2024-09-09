package reuse_test

import (
	"regexp"
	"testing"

	"github.com/valteem/reuse"
)

func TestCountMatches(t *testing.T) {

	tests := []struct {
		input   string
		pattern string
		count   int
	}{
		{"https://example.org:8080/article/101", "/", 4},
		{"https://example.org:8080?article=101?article=102", `\?`, 2}, // malformed URL query string
	}

	for _, tc := range tests {
		if count := reuse.CountMatches(tc.input, regexp.MustCompile(tc.pattern)); count != tc.count {
			t.Errorf("counting matches of %q in %q:\nget %d, expect %d", tc.pattern, tc.input, count, tc.count)
		}
	}
}
