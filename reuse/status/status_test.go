package status

import (
	"testing"
)

func TestState(t *testing.T) {

	tests := []struct {
		descr  string
		input  State
		output uint32
	}{
		{"NotReady", NotReady, 0},
		{"Ready", Ready, 1},
		{"Stopping", Stopping, 2},
		{"Idle", Idle, 3},
	}

	for _, tc := range tests {
		if output := uint32(tc.input); output != tc.output {
			t.Errorf("Status %q: get %d, expect %d", tc.descr, output, tc.output)
		}
	}

}
