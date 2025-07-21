package reuse_test

import (
	"testing"
)

// slicing [0:0] does not change slice capacity
func TestSlicingCap(t *testing.T) {

	s := []int{1, 2, 4, 8, 16, 32}

	capBeforeSlicing := cap(s)

	s = s[0:0]

	capAfterSlicing := cap(s)

	if capAfterSlicing != capBeforeSlicing {
		t.Errorf("expect same capacity before and after slicing, get %d and %d", capAfterSlicing, capBeforeSlicing)
	}

}
