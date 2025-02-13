package encode

import (
	"fmt"
	"testing"
)

func TestEncodeFloat64Bits(t *testing.T) {

	tests := []struct {
		input  float64
		output uint64
	}{
		{1.0, 4607182418800017408}, // 0 0(1)x10 (0)x52
		{0.1, 4591870180066957722},
		{1.7e+308, 9218378953502702454},
		{0.74e-323, 1},
	}

	for _, tc := range tests {
		if output := EncodeFloat64Bits(tc.input); output != tc.output {
			t.Errorf("encoding %s to bits: get %d, expect %d", fmt.Sprintf("%.2e", tc.input), output, tc.output)
		}
	}
}
