package reuse_test

import (
	"testing"
)

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
