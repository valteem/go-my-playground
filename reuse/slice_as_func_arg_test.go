package reuse_test

import (
	"testing"

	"github.com/valteem/reuse"
)

func TestAppendToSlice(t *testing.T) {

	s := []int{1, 2, 3, 4}
	reuse.AppendToSlice[int](s, 5)
	e := []int{1, 2, 3, 4} // nothind appended, slice is passed by value, which means length and capacity of 'outer' slice are not affected
	if !reuse.SliceEqual(s, e) {
		t.Errorf("Expected %+v, returned %+v", e, s)
	}

	s0 := []int{1, 2, 3, 4}
	reuse.AppendToSliceP[int](&s0, 5)
	e0 := []int{1, 2, 3, 4, 5}
	if !reuse.SliceEqual(s0, e0) {
		t.Errorf("Expected %+v, returned %+v", e0, s0)
	}

	c1 := reuse.NewSliceContainer[int]([]int{1, 2, 3, 4})
	c1.AppendToSlice(5)
	e1 := []int{1, 2, 3, 4} // nothind appended, function defined on SliceContainer struct, not pointer
	if !reuse.SliceEqual(c1.Slice(), e1) {
		t.Errorf("Expected %+v, returned %+v", e1, c1.Slice())
	}

	c2 := reuse.NewSliceContainer[int]([]int{1, 2, 3, 4})
	c2.AppendToSliceP(5)
	e2 := []int{1, 2, 3, 4, 5}
	if !reuse.SliceEqual(c2.Slice(), e2) {
		t.Errorf("Expected %+v, returned %+v", e2, c2.Slice())
	}

}

type wlist struct {
	elements []int
}

func TestSliceWrapper(t *testing.T) {
	w := &wlist{elements: []int{1, 2, 3}}
	f := func(e *[]int, v int) {
		*e = append(*e, v)

	}
	f(&w.elements, 4)
	if expectedlength, expectedvalue := len(w.elements), w.elements[3]; expectedlength != 4 || expectedvalue != 4 {
		t.Errorf("get slice length %d, expect 4, get value of last element %d, expect 4", expectedlength, expectedvalue)
	}

}
