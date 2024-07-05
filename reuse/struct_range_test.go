// https://stackoverflow.com/a/26166046

package reuse_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	reuse "github.com/valteem/reuse"
)

func TestSliceOfSquares(t *testing.T) {

	assert := assert.New(t)

	for _, tc := range []struct {
		input  []int32
		output []int32
	}{
		{
			[]int32{1, 2, 3},
			[]int32{1, 4, 9},
		},
		{
			input:  []int32{5, 6, 7},
			output: []int32{25, 36, 49},
		},
	} {
		result := reuse.SliceOfSquares(tc.input)
		assert.Equal(result, tc.output, "Shoud be equal")
	}

}
