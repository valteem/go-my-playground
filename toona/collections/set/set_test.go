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

func TestSetContains(t *testing.T) {

	e := []string{"one", "two", "three"}
	s := set.New(e...)
	
	if !s.Contains("three") {
		t.Errorf("set must contain element %s", "three")
	}
	if s.Contains("four") {
		t.Errorf("set does not contain element %s", "four")
	}

	e1 := []string{"a", "b", "c", "i", "j"}
	s1 := set.New(e1...)

	e2 := []string{"a", "c", "j"}
	s2 := set.New(e2...)
	if !s1.IsSuperSetOf(s2) {
		t.Error(`{"a", "b", "c", "i", "j"} is superset of {"a", "c", "j"}`)
	}

	e3 := []string{"a", "b", "k"}
	s3 := set.New(e3...)
	if s1.IsSuperSetOf(s3) {
		t.Error(`{"a", "b", "c", "i", "j"} is not superset of {"a", "b", "k"}`)
	}
}

func TestSetIntersection(t *testing.T) {

	e1 := []string{"a", "b", "c", "f", "g", "h"}
	s1 := set.New(e1...)
	e2 := []string{"b", "c", "h", "i", "j", "k"}
	s2 := set.New(e2...)

	s3 := s1.Intersection(s2)
	result := s3.ExtractSlice()
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	expected := []string{"b", "c", "h"}
	if !reflect.DeepEqual(result, expected) {
		t.Error("wrong intersection result")
	}

}