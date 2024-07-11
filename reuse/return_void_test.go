package reuse_test

import (
	"testing"
)

func getPair(x, xMin, xMax int) (output int, ok bool) {
	if x >= xMin && x <= xMax {
		return x, true
	}
	return // returns 0, false - `zero` values for corrresponding return types
}

type valuePair struct {
	x  int
	ok bool
}

func getValuePair(x, xMin, xMax int) (v *valuePair) {
	if x >= xMin && x <= xMax {
		return &valuePair{x: x, ok: true}
	}
	return // returns nil
}

func TestGetPair(t *testing.T) {
	tests := []struct {
		xInput    int
		xMinInput int
		xMaxInput int
		output    int
		ok        bool
	}{
		{1, 0, 2, 1, true},
		{1, 2, 3, 0, false},
	}
	for _, tc := range tests {
		if output, ok := getPair(tc.xInput, tc.xMinInput, tc.xMaxInput); output != tc.output || ok != tc.ok {
			t.Errorf("getPait(%d, %d, %d): get (%d, %t), expect (%d, %t)", tc.xInput, tc.xMinInput, tc.xMaxInput, output, ok, tc.output, tc.ok)
		}
	}
}

func TestGetValuePair(t *testing.T) {
	tests := []struct {
		xInput    int
		xMinInput int
		xMaxInput int
		output    int
		ok        bool
	}{
		{1, 0, 2, 1, true},
		{1, 2, 3, 0, false}, // panics here
	}
	for _, tc := range tests {
		v := getValuePair(tc.xInput, tc.xMinInput, tc.xMaxInput)
		if v == nil {
			panic("nil pointer")
		}
		if output, ok := v.x, v.ok; output != tc.output || ok != tc.ok {
			t.Errorf("getPait(%d, %d, %d): get (%d, %t), expect (%d, %t)", tc.xInput, tc.xMinInput, tc.xMaxInput, output, ok, tc.output, tc.ok)
		}
	}
}
