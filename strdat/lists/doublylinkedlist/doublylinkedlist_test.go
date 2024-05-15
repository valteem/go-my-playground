package doublylinkedlist

import (
	"testing"
)

func TestRemoveFirstLast(t *testing.T) {
	l := New[string]("a", "b", "c", "x")
	l.Remove(0)
	if l.first.prev != nil {
		t.Errorf("New list.first.prev should be nil")
	}
	l.Remove(2)
	if l.last.next != nil {
		t.Errorf("New list.last.next should be nil")
	}
	l.Remove(5)
	if l.size != 2 {
		t.Errorf("Trying to remove from index-out-of-bounds should have no effect")
	}
	l.Remove(0)
	if l.first.value != "c" || l.last.value != "c" || l.size != 1 {
		t.Errorf("List structure after three Remove() ops differs from expected")
	}
}
