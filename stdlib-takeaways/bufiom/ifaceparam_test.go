package bufiom_test

import (
	"lang.rev/stdlib-takeaways/bufiom"
	"testing"
)

func TestIfaceParam(t *testing.T) {

	var si bufiom.SomeSet = bufiom.SetOfInt{Item1: 1, Item2: 2}
	var ss bufiom.SomeSet = bufiom.SetOfStr{Item1: "1", Item2: "2"}

	bufiom.AssertSetType(si)
	bufiom.AssertSetType(ss)

	var sli bufiom.SomeSlice = bufiom.SliceOfInt{S: []int{1, 2, 3}}
	var sls bufiom.SomeSlice = bufiom.SliceOfStr{S: []string{"1", "2", "3"}}

	bufiom.AssertSliceType(sli)
	bufiom.AssertSliceType(sls)
	
}