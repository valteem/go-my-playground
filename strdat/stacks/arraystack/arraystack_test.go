package arraystack

import (
	"testing"
)

func TestPushPeekPop(t *testing.T) {
	input := []string{"apples", "pears", "cherries", "berries", "potatoes"}
	s := New[string]()

	for _, v := range input {
		s.Push(v)
	}

	for i := s.Size() - 1; i >= 0; i-- {
		if v, ok := s.Peek(); v != input[i] || ok != true {
			t.Errorf("Peek(): get (%s, %t), expect (%s, %t)", v, ok, input[i], true)
		}
		if size, empty := s.Size(), s.Empty(); size != (i+1) || empty {
			t.Errorf("Size(), Empty(): get %d, %t, expect %d, %t", size, empty, (i + 1), false)
		}
		if v, ok := s.Pop(); v != input[i] || ok != true {
			t.Errorf("Peek(): get (%s, %t), expect (%s, %t)", v, ok, input[i], true)
		}
	}
	if v, ok := s.Peek(); v != "" || ok {
		t.Errorf("Peek() on empty stack: get %s, %t, expect (%s), %t", v, ok, "", false)
	}
	if size, empty := s.Size(), s.Empty(); size != 0 || !empty {
		t.Errorf("Size(), Empty() on empty stack: get %d, %t, expect %d, %t", size, empty, 0, true)
	}
}
