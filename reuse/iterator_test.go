// https://go.dev/blog/range-functions
package reuse_test

import (
	"testing"
)

func TestBasicFuncRange(t *testing.T) {

	cap := 10

	f := func(g func(int) bool) {
		for i := 0; i <= cap; i++ {
			if !g(i) {
				return
			}
		}
	}

	i := 0
	for v := range f {
		if v != i {
			t.Errorf("get %d, expect %d", v, i)
		}
		i++
	}

}
