package linkedliststack

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

func TestStackIterator(t *testing.T) {
	input := []string{"apples", "pears", "cherries", "berries", "potatoes"}
	s := New[string]()

	it := s.Iterator()
	for it.Next() {
		t.Errorf("Cannot iterate over empty stack")
	}

	for _, v := range input {
		s.Push(v)
	}

	it.Begin()
	if actualIndex, expectedIndex := it.Index(), -1; actualIndex != expectedIndex {
		t.Errorf("Index() after Begin(): get %d, expect %d", actualIndex, expectedIndex)
	}

	it.First()
	if actualIndex, expectedIndex := it.Index(), 0; actualIndex != expectedIndex {
		t.Errorf("Index() after First(): get %d, expect %d", actualIndex, expectedIndex)
	}
	if actualValue, expectedValue := it.Value(), "potatoes"; actualValue != expectedValue {
		t.Errorf("Value() after First(): get %s, expect %s", actualValue, expectedValue)
	}
	it.Next()
	if actualIndex, expectedIndex := it.Index(), 1; actualIndex != expectedIndex {
		t.Errorf("Index() after Next(): get %d, expect %d", actualIndex, expectedIndex)
	}
	if actualValue, expectedValue := it.Value(), "berries"; actualValue != expectedValue {
		t.Errorf("Value() after NextTo(): get %s, expect %s", actualValue, expectedValue)
	}

	it.NextTo(func(i int, v string) bool {
		return v == "pears"
	})
	if actualIndex, expectedIndex := it.Index(), 3; actualIndex != expectedIndex {
		t.Errorf("Index() after NextTo(): get %d, expect %d", actualIndex, expectedIndex)
	}
	if actualValue, expectedValue := it.Value(), "pears"; actualValue != expectedValue {
		t.Errorf("Value() after NextTo(): get %s, expect %s", actualValue, expectedValue)
	}

	it.NextTo(func(i int, v string) bool {
		return v == "onions"
	})
	if actualIndex, expectedIndex := it.Index(), 5; actualIndex != expectedIndex {
		t.Errorf("Index() after NextTo(): get %d, expect %d", actualIndex, expectedIndex)
	}
	if actualValue, expectedValue := it.Value(), ""; actualValue != expectedValue {
		t.Errorf("Value() after NextTo(): get %s, expect %s", actualValue, expectedValue)
	}
}
