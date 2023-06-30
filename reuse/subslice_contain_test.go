package reuse_test

import (
	"fmt"
	"testing"
	"github.com/valteem/reuse"
)

func TestSliceContains(t *testing.T) {

	s1 := []int{1, 2, 3, 4}
	s2 := []int{1, 2, 3, 4}
	s3 := []int{1, 2, 3}
	s4 := []int{1, 2, 3, 5}

	fmt.Println(reuse.SliceEqual(s1, s2), reuse.SliceEqual(s1, s3), reuse.SliceEqual(s1, s4))

	a := []int{1, 2, 4, 7, 10, 8, 11, 1, 2, 4, 7}
	b := []int{1, 2, 4}
	fmt.Println(reuse.SliceContains(a, b))
	
}