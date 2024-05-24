package reuse_test

import (
	"fmt"
	"testing"
)

type labelName string

func TestSingleQuoteFormat(t *testing.T) {
	tests := []struct {
		ln     labelName
		output string
	}{
		{
			ln:     "some label",
			output: "\"some label\"",
		},
		{
			ln:     "some other label with numbers 1234",
			output: "\"some other label with numbers 1234\"",
		},
	}
	for _, tst := range tests {
		if s := fmt.Sprintf("%q", tst.ln); s != tst.output {
			t.Errorf("Get %s, expect %s", s, tst.output)
		}
	}
}
