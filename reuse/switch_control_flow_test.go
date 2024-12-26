package reuse_test

import (
	"slices"
	"testing"
)

func switchNumberRange(input int) []string {
	output := []string{}
	switch {
	case input != 0:
		output = append(output, "non-zero")
	case input > 0:
		output = append(output, "positive")
	case (input/2)*2 == input:
		output = append(output, "even")
	case (input/2)*2 < input:
		output = append(output, "uneven")
	}
	return output
}

func TestSwitchNumberRange(t *testing.T) {

	tests := []struct {
		input  int
		output []string
	}{
		{1, []string{"non-zero"}}, // only first matching case is executed
	}

	for _, tc := range tests {
		if output := switchNumberRange(tc.input); !slices.Equal(output, tc.output) {
			t.Errorf("get\n%v\nexpect\n%v\n", output, tc.output)
		}
	}

}
