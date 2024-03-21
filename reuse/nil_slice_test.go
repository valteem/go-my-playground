// https://stackoverflow.com/a/44305910

package reuse_test

import (
	"fmt"
	"testing"
)

func TestInitNilSlice(t *testing.T) {

	in := []string(nil) // most probably useful for inline nil slice declaration
	// impossible condition: nil != nil
	// if in != nil {
	// 	t.Errorf("slice must be nil")
	// }
	add := []string{"first, second", "third", "closing"}
	out := append(in, add...)
	for i, v := range out {
		if v != add[i] {
			t.Errorf("two elements must be equal: %v and %v", v, add[i])
		}
	}
}

func TestSliceReassign(t *testing.T) {
	s := []string{"first", "second"}
	c := s
	c[0] = "start"
	if s[0] != "start" {
		t.Errorf("front slice element is %s, expected %s", s[0], c[0])
	}
	// both slice arrays are the same
	cFirstAddr := fmt.Sprintf("%p", &c[0])
	sFirstAddr := fmt.Sprintf("%p", &s[0])
	if &cFirstAddr != &sFirstAddr {
		t.Errorf("pointer to `copied` slice front element is %s, expected %s", cFirstAddr, sFirstAddr)
	}
	// however other slice params are not
	cAddr := fmt.Sprintf("%p", &c)
	sAddr := fmt.Sprintf("%p", &s)
	if cAddr == sAddr {
		t.Errorf("pointer to `copied` slice is %s", cAddr)
	}
}

func TestShallowSliceCopyCapLen(t *testing.T) {
	s := []int{-2, -1}
	c := s
	c[0] = 1
	c[1] = 2
	for i := 3; i <= 128; i++ {
		c = append(c, i)
	}
	if s[0] != c[0] {
		t.Errorf("%d %d", s[0], c[0])
	}
	if cap(s) == cap(c) {
		t.Errorf("%d %d", cap(s), cap(c))
	}
}
