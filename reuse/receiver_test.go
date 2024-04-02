package reuse_test

import (
	"fmt"
	"testing"
)

// Basic example: value vs pointer receivers for methods chnaging wrapped objects

type Wrapper[T any] struct {
	values []T
}

func (w Wrapper[T]) ClearV() {
	w.values = []T{} //lint:ignore SA4005 for demonstration purposes
}

func (w *Wrapper[T]) ClearP() {
	w.values = []T{}
}

func TestClear(t *testing.T) {

	wValue := Wrapper[int]{
		values: []int{1, 2, 3},
	}
	wValue.ClearV()
	if l := len(wValue.values); l != 3 {
		t.Errorf("Wrapper Clear by Value: expect length = 0, get %d", l)
	}

	wPoint := Wrapper[int]{
		values: []int{1, 2, 3},
	}
	wPoint.ClearP()
	if l := len(wValue.values); l != 3 {
		t.Errorf("Wrapper Clear by Pointer: expect length = 0, get %d", l)
	}
}

// Example: how to make pointer receiver comply with interface type

type container interface {
	clear()
	show() string
}

// Does nothing, used to test if an argument complies with `container` interface
func justComply(c container) string {
	return c.show()
}

type mycont struct {
	i int
}

func (m *mycont) clear() {
	m.i = 0
}

func (m mycont) show() string {
	return fmt.Sprintf("%d", m.i)
}

func TestComply(t *testing.T) {

	// This does not compile:

	//	cv := mycont{}
	//	justComply(cv)

	// cannot use cv (variable of type mycont) as container value in argument to justComply:
	// mycont does not implement container (method clear has pointer receiver)

	// This compiles:
	// using pointer instead of value as function argument to comply with `container` interface
	cp := &mycont{1}
	if s := justComply(cp); s != "1" {
		t.Errorf("expect 1, get %s", s)
	}

}
