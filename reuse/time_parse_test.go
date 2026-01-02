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

func TestDurationFormat(t *testing.T) {

	tests := []struct {
		input  time.Duration
		output string
	}{
		{
			input:  time.Duration(11*60*60+21*60+31) * 1000000000,
			output: "11h21m31s",
		},
		{
			input:  time.Duration(1*60*60+21) * 1000000000,
			output: "1h0m21s",
		},
		{
			input:  time.Duration(1) * 100000000, // 10^8
			output: "100ms",
		},
		{
			input:  time.Duration(1) * 100000, // 10^5
			output: "100µs",
		},
		{
			input:  time.Duration(1) * 100,
			output: "100ns",
		},
	}

	for _, tc := range tests {
		if output := tc.input.String(); output != tc.output {
			t.Errorf("string representation of duration input: get %q, expect %q", output, tc.output)
		}
	}
}
