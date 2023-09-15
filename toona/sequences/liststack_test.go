package sequences_test

import (
	"testing"

	"my.play.go/toona/sequences"

)

func TestListStack(t *testing.T) {

	s := sequences.NewListStack[int]()

	if s.IsEmpty() != true {
		t.Error("wrong IsEmpty() result: should be true")
	}

	for i := 1; i <= 10; i++ {
		s.Push(i)
	}

	for i := 10; i >= 1; i-- {
		v, _ := s.Pop()
		if v != i {
			t.Errorf("wrong Pop() result: expect %v, get %v", i, v)
		}
	}

	v, e := s.Pop()
	if e.Error() != "stack empty" || v != 0 {
		t.Errorf("wrong Pop() result from empty stack: %v %v", v, e)
	}
}