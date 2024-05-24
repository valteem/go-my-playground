package reuse_test

import (
	"testing"
)

func TestRangeOverString(t *testing.T) {
	tests := []struct {
		desc       string
		testString string
	}{
		{
			desc:       "lowercase letters only",
			testString: "abcuvw",
		},
		{
			desc:       "numbers only",
			testString: "01234567",
		},
		{
			desc:       "includes capital letters",
			testString: "0A2B4C7u",
		},
	}
	for _, tst := range tests {
		for i, b := range tst.testString {
			if !((b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')) {
				t.Errorf("%s: get %v at position %d", tst.desc, b, i)
			}
		}
	}
}
