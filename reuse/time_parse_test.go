package reuse_test

import (
	"time"

	"testing"
)

const (
	layout        = "2006-01-02 15:04:05-0700" // layout must use reference values for year, month, day etc.
	offsetSeconds = int(4 * 60 * 60)           // 4 hours
)

func TestTimeParse(t *testing.T) {

	tests := []struct {
		input  string
		output time.Time
	}{
		{
			input:  "2025-12-30 19:30:00+0400",
			output: time.Date(2025, time.Month(12), 30, 19, 30, 0, 0, time.FixedZone("UTC+4", offsetSeconds)),
		},
	}

	for _, tc := range tests {
		output, err := time.Parse(layout, tc.input)
		if err != nil {
			t.Errorf("failed to parse input %q: %v", tc.input, err)
			continue
		}
		if !output.Equal(tc.output) {
			t.Errorf("get %v for %q input, expect %v", output, tc.input, tc.output)
		}
	}
}
