package reuse_test

import (
	"reflect"
	"regexp"

	"testing"
)

var (
	re    = regexp.MustCompile(`(?P<name>.*) (?P<occupation>.*)`)
	names = re.SubexpNames()
)

func TestCaptureGroup(t *testing.T) {

	tests := []struct {
		input        string
		captureGroup string
		output       string
	}{
		{
			input:        `Pham driver`,
			captureGroup: "name",
			output:       "Pham",
		},
		{
			input:        `Pham driver`,
			captureGroup: "occupation",
			output:       "driver",
		},
	}

	for _, tc := range tests {
		match := re.FindStringSubmatch(tc.input)

		for i, n := range names {
			if (i > 0) && n == tc.captureGroup && match[i] != tc.output {
				t.Errorf("capture group %q^ get %q, expect %q", tc.captureGroup, match[i], tc.output)
			}
		}
	}
}

func TestMultipleCapture(t *testing.T) {

	tests := []struct {
		pattern string
		input   string
		output  map[string]string
	}{
		{
			pattern: `(?P<num1>\d+) (?P<text1>\D+) (?P<num2>\d+) (?P<text2>.*)`,
			input:   `123 onions 456 apples`,
			output:  map[string]string{"num1": "123", "text1": "onions", "num2": "456", "text2": "apples"},
		},
	}

	for _, tc := range tests {
		re = regexp.MustCompile(tc.pattern)
		matches := re.FindStringSubmatch(tc.input) // captures nothing if input does not fully match pattern
		names := re.SubexpNames()
		captures := map[string]string{}
		for i, n := range names {
			if i > 0 {
				captures[n] = matches[i]
			}
		}
		if !reflect.DeepEqual(captures, tc.output) {
			t.Errorf("get\n%v\nexpect\n%v\n", captures, tc.output)
		}
	}

}
