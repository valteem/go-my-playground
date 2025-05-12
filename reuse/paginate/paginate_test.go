package paginate

import (
	"slices"
	"testing"
)

func TestPaginate(t *testing.T) {

	tests := []struct {
		spec   string
		input  []int
		start  int
		size   int
		output []int
	}{
		{"regular pagination", []int{0, 1, 2, 3, 4}, 2, 2, []int{2, 3}},
		{"end beyond input length", []int{0, 1, 2, 3, 4}, 3, 2, []int{4}},
		{"start beyond input length", []int{0, 1, 2, 3, 4, 5}, 4, 2, []int{}},
	}

	for _, tc := range tests {
		if output := Paginate(tc.input, tc.start, tc.size); !slices.Equal(output, tc.output) {
			t.Errorf("%s:\nget\n%v\nexpect\n%v\n", tc.spec, output, tc.output)
		}
	}

}
