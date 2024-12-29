package reuse_test

import (
	"strconv"
	"testing"
)

func TestParseInt(t *testing.T) {

	tests := []struct {
		input   string
		bitSize int
		output  int64
	}{
		{"256", 8, 127},
		{"65636", 16, 32767}, // one bit for sign
		{"128", 32, 128},
		{"128", 64, 128},
	}

	for _, tc := range tests {
		output, _ := strconv.ParseInt(tc.input, 10, tc.bitSize)
		if output != tc.output {
			t.Errorf("%v: get %d, expect %d", tc, output, tc.output)
		}
	}

}
