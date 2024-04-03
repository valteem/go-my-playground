package reuse_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestListClear(t *testing.T) {
	l := []int{0, 1, 2, 3, 4}
	index := 2
	clear(l[index : index+1])
	expect := []int{0, 1, 0, 3, 4} // clear() does not remove an element, just sets it to `initial` value for list elements type
	if !cmp.Equal(l, expect) {
		t.Errorf("get %v, expect %v", l, expect)
	}
}
