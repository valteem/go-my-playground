package set_test

import (
	"reflect"
	"sort"
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
	for k := range s {
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

	r := []string{"two", "three"}
	s.RemoveMany(r...)
	for _, v := range r {
		if _, in := s[v]; in {
			t.Errorf("error rremoving element %s", v)
		}
	
	}
	
}

func TestSetMerge(t *testing.T) {

	e1 := []string{"one", "two", "three"}
	s1 := set.New(e1...) // generic argument type inference

	e2 := []string{"three", "four", "five"}
	s2 := set.New(e2...)

	s1.Merge(s2)

	expected := []string{"one", "two", "three", "four", "five"}
	result := s1.ExtractSlice()
	//TODO: replace inline sorting with custom slice sort
	sort.Slice(expected, func(i, j int) bool {
		return expected[i] < expected[j]
	})
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	if !reflect.DeepEqual(expected, result) {
		t.Error("error merging slices")
	}
}