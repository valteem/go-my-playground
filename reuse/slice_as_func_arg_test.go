package reuse_test

import (
	"fmt"
	"testing"

	"github.com/valteem/reuse"
)

func TestAppendToSlice(t *testing.T) {
	
	s := []int{1, 2, 3, 4}
	reuse.AppendToSlice[int](s, 5)
	fmt.Println(s) // [1 2 3 4]

	c1 := reuse.NewSliceContainer[int]([]int{1, 2, 3, 4})
	c1.AppendToSlice(5)
	fmt.Println(c1.Slice()) // [1 2 3 4]

	c2 := reuse.NewSliceContainer[int]([]int{1, 2, 3, 4})
	c2.AppendToSliceP(5)
	fmt.Println(c2.Slice()) // [1 2 3 4 5]

}