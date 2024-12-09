package reuse_test

import (
	"errors"
	"testing"

	"github.com/valteem/reuse"
)

func TestErrorTypes(t *testing.T) {

	tests := []struct {
		err     error
		errMsg  string
		errCode int
	}{
		{&reuse.SignError{}, reuse.SignErrorMsg, 1},
		{&reuse.RangeError{LowerBound: 1, UpperBound: 2}, "value should be between 1 and 2", 2},
		{errors.New("some other error"), "some other error", -1},
	}

	var errCode int
	for _, tc := range tests {
		switch tc.err.(type) {
		case *reuse.SignError:
			errCode = 1
		case *reuse.RangeError:
			errCode = 2
		default:
			errCode = -1
		}

		if errMsg := tc.err.Error(); errMsg != tc.errMsg || errCode != tc.errCode {
			t.Errorf("error message/code: get %q/%d, expect %q/%d", errMsg, errCode, tc.errMsg, tc.errCode)
		}

	}

}
