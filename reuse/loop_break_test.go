package reuse_test

import (
	"testing"
)

func loopBreakSingle() int {
	output := 0
Loop:
	for i := 0; i <= 5; i++ {
		if i >= 5 {
			break Loop
		}
		output++

	}
	return output
}

func loopBreakInnerContinueOuter() int {
	output := 0
LoopOuter:
	for i := 0; i <= 5; i++ {
		if i > 1 {
			continue LoopOuter
		}
		output++
	LoopInner:
		for j := 0; j <= 5; j++ {
			if j > 1 {
				break LoopInner
			}
			output++
		}
	}
	return output
}
func TestLoopBreak(t *testing.T) {

	tests := []struct {
		name   string
		f      func() int
		output int
	}{
		{"single loop with break", loopBreakSingle, 5},
		{"double loop, break inner, continue outer", loopBreakInnerContinueOuter, 6},
	}

	for _, tc := range tests {
		output := tc.f()
		if output != tc.output {
			t.Errorf("%s:\nget %d, expect %d", tc.name, output, tc.output)
		}
	}

}
