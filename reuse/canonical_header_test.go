package reuse_test

import (
	"net/http"
	"testing"
)

func TestCanonicalHeaderKeys(t *testing.T) {

	tests := []struct {
		input  string
		output string
	}{
		{"x-sidecar-response", "X-Sidecar-Response"},
	}

	for _, tc := range tests {
		if output := http.CanonicalHeaderKey(tc.input); output != tc.output {
			t.Errorf("input %s: get %s, expect %s", tc.input, output, tc.output)
		}
	}
}
