package bag

import (
//	"fmt"
	"testing"
)

func TestBag(t *testing.T) {

	b := New[string]()
	b.Insert("item1")
	b.Insert("item1")
	if b.Count("item1") != 2 {
		t.Errorf("wrong item count")
	}
	b.Reset()
	if b.size != 0 {
		t.Errorf("wrong bag size after reset")
	}

}