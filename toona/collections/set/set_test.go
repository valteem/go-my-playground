package set_test

import (
	"testing"

	"my.play.go/toona/collections/set"
)

func TestSetNew(t *testing.T) {

	e := []string{"one", "two", "three"}
	s := set.New[string](e...)
	if len(s) != len(e) {
		t.Errorf("error set initialization - invalid length")
	}

// all slice elements are in map	
	for _, v := range e {
		if _, in := s[v]; !in {
			t.Errorf("error set initialization - element not found %s", v)
		}
	}

// all map keys are in slice
	for k, _ := range s {
		in := false
		for _, v := range e {
			if k == v {
				in = true
			}
		}
		if !in {
			t.Errorf("error set initialization - redundant elements in map %s", k)
		}
	}
}


