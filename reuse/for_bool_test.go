package reuse_test

import (
	"testing"
)

func TestForBool(t *testing.T) {
	a, b, c := 1, 3, 0
	for a < b {
		c++
		a++
	}
	if c != 2 {
		t.Errorf("get %d, expect 2", c)
	}
}
