package errors

import (
	"fmt"
	"testing"
)

func TestCanonicalCode(t *testing.T) {

	tests := []struct {
		description string
		err         error
		expected    Code
	}{
		{
			description: "stdlib error",
			err:         fmt.Errorf("stdlib error"),
			expected:    Unspecified,
		},
		{
			description: "custom error with NotFound code",
			err: &Error{
				Code: NotFound,
			},
			expected: NotFound,
		},
		{
			description: "custom error with Internal nested error",
			err: &Error{
				Code: Unspecified,
				Err: &Error{
					Code: Internal,
				},
			},
			expected: Internal,
		},
	}

	for _, tst := range tests {
		c := CanonicalCode(tst.err)
		if c != tst.expected {
			t.Errorf("expect %v error code, get %v instead", tst.expected, c)
		}
	}

}
