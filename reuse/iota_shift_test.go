package reuse_test

import (
	"testing"
)

const (
	c0 = 1 << iota
	c1
	c2
	c3 = iota
	c4 = 1 << iota
	c5 = iota
)

func TestIotaShift(t *testing.T) {

	tests := []struct {
		input  int
		output int
	}{
		{c0, 1 << 0},
		{c1, 1 << 1},
		{c2, 1 << 2},
		{c3, 3},
		{c4, 1 << 4},
		{c5, 5},
	}

	for _, tc := range tests {
		if tc.input != tc.output {
			t.Errorf("get %d, expect %d", tc.input, tc.output)
		}
	}
}
