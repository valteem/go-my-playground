package reuse_test

import (
	"encoding/binary"
	"testing"
)

func TestBigLittleEndian(t *testing.T) {

	tests := []struct {
		input  uint16
		output uint16
	}{
		{1 << (7 + 8), 1 << 7},
		{1 << (6 + 8), 1 << 6},
		{1 << (5 + 8), 1 << 5},
		{1 << (4 + 8), 1 << 4},
		{1 << (3 + 8), 1 << 3},
		{1 << (2 + 8), 1 << 2},
		{1 << (1 + 8), 1 << 1},
		{1 << 8, 1},
		{1 << 7, 1 << (7 + 8)},
		{1 << 6, 1 << (6 + 8)},
		{1 << 5, 1 << (5 + 8)},
		{1 << 4, 1 << (4 + 8)},
		{1 << 3, 1 << (3 + 8)},
		{1 << 2, 1 << (2 + 8)},
		{1 << 1, 1 << (1 + 8)},
		{1 << 0, 1 << (0 + 8)},
	}

	for _, tc := range tests {
		b := make([]byte, 2)
		binary.BigEndian.PutUint16(b, tc.input)
		output := binary.LittleEndian.Uint16(b)
		if output != tc.output {
			t.Errorf("BigEndian->LittleEndian: get %d, expect %d", output, tc.output)
		}
	}
}
