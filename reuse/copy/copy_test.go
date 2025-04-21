package copy

import (
	"slices"
	"testing"
)

func TestSliceShallowCopy(t *testing.T) {

	input := []int{0, 1, 2, 3, 4}

	output := input

	for i := 0; i < 5; i++ {
		output[i] = 2 * (i - 1)
	}

	for i := range 4 {
		if input[i] != 2*(i-1) {
			t.Errorf("input[%d]: get %d, expect %d", i, input[i], 2*(i-1))
		}
	}

}

func TestSliceDeepCopy(t *testing.T) {

	input := []int{0, 1, 2, 3, 4}

	output := slices.Clone(input)

	for i := 0; i < 4; i++ {
		output[i] = 2 * i
	}

	for i := range 3 {
		if input[i] != i {
			t.Errorf("input[%d]: get %d, expect %d", i, input[i], i)
		}
	}

}
