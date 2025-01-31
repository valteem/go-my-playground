package lists

import (
	"testing"
)

func TestCircularList(t *testing.T) {

	cl := NewCircularList[int]()
	listLen := 10
	elems := make([]*ListElement[int], 0, listLen)

	for i := 0; i < listLen; i++ {
		cl.Add(i)
		elems = append(elems, cl.Last())
	}

	a := cl.anchor
	var nextExpected, prevExpected *ListElement[int]
	for i := 0; i < listLen; i++ {
		if i == 0 {
			nextExpected, prevExpected = a, elems[i+1]
		} else if i == listLen-1 {
			nextExpected, prevExpected = elems[i-1], a
		} else {
			nextExpected, prevExpected = elems[i-1], elems[i+1]
		}
		if elems[i].next != nextExpected || elems[i].prev != prevExpected {
			t.Errorf("next: get %p, expect %p; prev: get %p, expect %p\n", elems[i].next, nextExpected, elems[i].prev, prevExpected)
		}
	}

}
