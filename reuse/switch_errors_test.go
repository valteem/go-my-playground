package reuse_test

import (
	"errors"
	"fmt"
	"testing"
)

type errNested struct {
	input string
}

func (e errNested) Error() string {
	return fmt.Sprintf("invalid input %s", e.input)
}

var (
	errNotFound = fmt.Errorf("not found")
	errUnknown  = fmt.Errorf(("unknown"))
	errWrapper  = fmt.Errorf("wrapper around %w", errNested{input: "abc"})
)

func TestCustomErrors(t *testing.T) {

	tests := []struct {
		input  error
		output string
	}{
		{errNotFound, "not found"},
		{errUnknown, "unknown"},
		{errWrapper, "nested"},
		{fmt.Errorf("other"), "other"},
	}

	for _, tc := range tests {
		// Use switch instead of Error() for testing purposes only
		var output string
		var nested errNested
		switch {
		case errors.Is(tc.input, errNotFound):
			output = "not found"
		case errors.Is(tc.input, errUnknown):
			output = "unknown"
		case errors.As(tc.input, &nested):
			output = "nested"
		default:
			output = "other"
		}
		if output != tc.output {
			t.Errorf("%v: get %q, expect %q", tc.input, output, tc.output)
		}
	}
}
