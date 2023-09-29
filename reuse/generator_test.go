package reuse_test

import (
	"testing"

	"github.com/valteem/reuse"
)

func TestGenerator(t *testing.T) {

	c := reuse.NewContainerIterable[int]()

	for i := 0; i < 5; i++ {
		c.Add(i)
	}

	r := make([]int, 0, 5)
	for v := range c.Iterate() {
		r = append(r, v)
	}

	for i := 0; i < 5; i++ {
		if r[i] != i {
			t.Errorf("wrong Iterate() result: expected %+v, received %+v", i, r[i])
		}
	}
}