// Always prefer errors.Is and errors.As over type assertion
// https://www.reddit.com/r/golang/comments/1bwagd1/beginner_should_i_use_errorsas_or_type_assertion/?rdt=47398

package errortypeassert

import (
	"testing"
)

func TestErrorTypeAssert(t *testing.T) {

	tests := []struct {
		input     error
		assert    AssertErrorType
		outputErr error
		outputOk  bool
	}{
		{
			input: ErrStr,
			assert: func(e error) (error, bool) {
				output, ok := e.(*StrError)
				return output, ok
			},
			outputErr: ErrStr,
			outputOk:  true,
		},
		{
			input: ErrStr,
			assert: func(e error) (error, bool) {
				output, ok := e.(*IntError)
				return output, ok
			},
			outputErr: (*IntError)(nil), // not just 'nil'
			outputOk:  false,
		},
	}

	for _, tc := range tests {
		outputErr, outputOk := tc.assert(tc.input)
		if outputErr != tc.outputErr || outputOk != tc.outputOk {
			t.Errorf("asserting %v:\nget\n%v / %t\nexpect\n%v, %t\n",
				tc.input,
				outputErr,
				outputOk,
				tc.outputErr,
				tc.outputOk)
		}
	}

}
