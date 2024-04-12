package singlylinkedlist

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRemove(t *testing.T) {

	list := New[int](0, 1, 2, 3, 4)

	list.Remove(0)
	if v, b := list.Get(0); v != 1 || b != true {
		t.Errorf("Get (%v, %t) at the head, expect (1, true", v, b)
	}

}

func TestValues(t *testing.T) {

	list := New[int]()
	expect := []int{}
	for i := 0; i < 1<<16; i++ {
		list.Add(i)
		expect = append(expect, i)
	}
	if v := list.Values(); !cmp.Equal(v, expect) {
		t.Errorf("Get %v, expect %v", v, expect)
	}
}
