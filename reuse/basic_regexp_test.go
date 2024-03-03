package reuse_test

import (
	"regexp"
	"testing"
)

const (
	regexDate  = `[0-9]{4}/[0-9]{2}/[0-9]{2}`
	regexMicro = `\.[0-9]{6}`
)

func TestBasicRegexp(t *testing.T) {

	tests := []struct {
		input   string
		pattern string
	}{
		{
			input:   "2024/03/03",
			pattern: regexDate,
		},
		{
			input:   ".123456",
			pattern: regexMicro,
		},
	}

	for _, tst := range tests {
		match, e := regexp.MatchString(tst.pattern, tst.input)
		if e != nil {
			t.Fatal("pattern does not compile:", e)
		}
		if !match {
			t.Errorf("string %q does not match pattern %q", tst.input, tst.pattern)
		}
	}

}
