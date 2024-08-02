package container

import (
	"slices"
	"testing"

	"golang.org/x/tools/container/intsets"
)

const (
	sizeSet = 10
)

func TestSparsebasic(t *testing.T) {

	set := &intsets.Sparse{}
	feedTo := []int{}
	feedToExpected := []int{}
	for i := range sizeSet {
		set.Insert(i)
		feedToExpected = append(feedToExpected, i)
	}

	for i := range sizeSet {
		if !set.Has(i) {
			t.Errorf("expect %d to be in the set", i)
		}
	}

	feedTo = set.AppendTo(feedTo) // need to specify feedTo as both argument and receiver
	if !slices.Equal(feedTo, feedToExpected) {
		t.Errorf("get %v, expect %v", feedTo, feedToExpected)
	}

}
