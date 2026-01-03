package reuse_test

import (
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
