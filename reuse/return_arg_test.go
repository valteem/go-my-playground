package reuse_test

import (
	"fmt"
	"slices"
	"testing"
)

func changeNoEffect(value int) {
	value += 1
}

// change argument value without passing pointer as argument
func changeEffect(value int) int {
	value += 1
	return value
}

func changeSliceAndReturn(s []int) []int {
	s = append(s, 1)
	return s
}

func TestReturnArg(t *testing.T) {

	v1 := 1
	changeNoEffect(v1)
	if v1 != 1 {
		t.Errorf("changeNoEffect(): get %d, expect 1", v1)
	}

	v2 := 1
	v2 = changeEffect(v2)
	if v2 != 2 {
		t.Errorf("changeEffect(): get %d, expect 2", v2)
	}

}

func TestReturnSlice(t *testing.T) {

	s := []int{0}
	pAddrBefore := fmt.Sprintf("%p", &s[0])
	s = changeSliceAndReturn(s) // returns new slice
	pAddrAfter := fmt.Sprintf("%p", &s[0])
	if !slices.Equal(s, []int{0, 1}) {
		t.Errorf("changeSliceAndReturn(): get %v, expect [0, 1]", s)
	}
	if pAddrAfter == pAddrBefore {
		t.Errorf("changeSliceAndReturn(): expect new slice allocation, get new pointer %s and old one %s", pAddrAfter, pAddrBefore)
	}

	// apply same steps to append()
	a := []int{0}
	aAddrBefore := fmt.Sprintf("%p", &s[0])
	a = append(a, 1) // returns same slice
	aAddrAfter := fmt.Sprintf("%p", &s[0])
	if aAddrAfter != aAddrBefore {
		t.Errorf("append(): expect no new slice allocation, get new pointer %s and old one %s", aAddrAfter, aAddrBefore)
	}

}
