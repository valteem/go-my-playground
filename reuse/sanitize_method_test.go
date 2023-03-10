package reuse_test

import (
	"testing"
	"github.com/valteem/reuse"
)

func TestSanitizeMethod(t *testing.T) {
	tests := []struct {
		input string
		expected string
	}{
		{input: "GET", expected: "GET"},
		{input: "get", expected: "GET"},
		{input: "POST", expected: "POST"},
		{input: "post", expected: "POST"},
		{input: "9999", expected: "OTHER"},
	}

	for _, out := range tests {
		if reuse.SanitizeMethod(out.input) != out.expected {
			t.Errorf("%v output should be equal %v", out.input, out.expected)
		}
	}

}