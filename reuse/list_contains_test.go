package reuse_test

import (
	"fmt"
	"testing"
	"github.com/valteem/reuse"
)

func comp(a int, b int) bool {
	if a==b {
		return true
	} else {
		return false
	}
}

func TestListContains(t *testing.T) {
	l := reuse.List[int]{1, 2, 3}
	fmt.Println(l.Contains(2, comp))
	fmt.Println(l.Contains(4, comp))
}