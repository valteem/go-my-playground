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

func TestListIterator(t *testing.T) {
	input := []string{"apples", "pears", "cherries", "berries", "potatoes"}
	s := New[string]()
	it := s.Iterator()
	for it.Next() {
		t.Errorf("Cannot iterate over empty stack - Next()")
	}
	for it.Prev() {
		t.Errorf("Cannot iterate over empty stack - Prev()")
	}
	for _, v := range input {
		s.Push(v)
	}

	l := len(input)
	for it.Next() {
		if i, v := it.Index(), it.Value(); v != input[l-i-1] { // begin with "potatoes" - top of the stack
			t.Errorf("Value() after Next(): get %s, expect %s", v, input[l-i-1])
		}
	}
	for it.Prev() {
		if i, v := it.Index(), it.Value(); v != input[l-i-1] { // begin with "apples" - bottom of the stack
			t.Errorf("Value() after Prev(): get %s, expect %s", v, input[l-i-1])
		}
	}
}

func TestNextToPrevTo(t *testing.T) {
	input := []string{"apples", "pears", "cherries", "berries", "potatoes"}
	s := New[string]()
	for _, v := range input {
		s.Push(v)
	}
	it := s.Iterator()
	it.NextTo(func(i int, v string) bool {
		return v == "pears"
	})
	if i, v := it.Index(), it.Value(); i != 3 || v != "pears" {
		t.Errorf("NextTo(): get (%d, %s), expect (%d, %s)", i, v, 3, "pears")
	}
	it.PrevTo(func(i int, v string) bool {
		return i == 1
	})
	if i, v := it.Index(), it.Value(); i != 1 || v != "berries" {
		t.Errorf("NextTo(): get (%d, %s), expect (%d, %s)", i, v, 1, "berries")
	}
	it.NextTo(func(i int, v string) bool {
		return v == "onions"
	})
	if i, v := it.Index(), it.Value(); i != 5 || v != "" {
		t.Errorf("NextTo(): get (%d, %s), expect (%d, %s)", i, v, 5, "")
	}
	it.PrevTo(func(i int, v string) bool {
		return i == 11
	})
	if i, v := it.Index(), it.Value(); i != -1 || v != "" {
		t.Errorf("NextTo(): get (%d, %s), expect (%d, %s)", i, v, -1, "")
	}
}
