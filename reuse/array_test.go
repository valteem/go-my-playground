// https://victoriametrics.com/blog/go-array/
package reuse_test

import (
	"reflect"
	"slices"
	"testing"
)

func TestArrayForRange(t *testing.T) {

	a, b := [4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}

	var outputActual []int
	for i, v := range a {
		if i == 1 {
			a = b // this value of b is never used
		}
		// v is actually value from a hidden copy of 'a', not 'a' itself
		outputActual = append(outputActual, v)
	}

	outputExpected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(outputActual, outputExpected) {
		t.Errorf("output:\nget\n%v\nexpect\n%v", outputActual, outputExpected)
	}

	// Comparison operator '==' can be applied to arrays, no need for reflect
	if a != b {
		t.Errorf("expect two arrays hold same values, get:\n%v\n%v", a, b)
	}

	clear(outputActual) // clears to [0, 0, 0, 0]

	var outputActualWithPointer []int
	ap := [4]int{1, 2, 3, 4}
	for i, v := range &ap {
		if i == 1 {
			ap = b
		}
		outputActualWithPointer = append(outputActualWithPointer, v)
	}
	outputExpectedWithPointer := []int{1, 2, 7, 8}
	if !reflect.DeepEqual(outputActualWithPointer, outputExpectedWithPointer) {
		t.Errorf("output for range with pointer:\nget\n%v\nexpect\n%v\n", outputActualWithPointer, outputExpectedWithPointer)
	}

}

func TestArraySlicing(t *testing.T) {

	a := [4]byte{1, 2, 4, 8}
	b := a[:] // []byte
	if actual, expected := len(b), 4; actual != expected {
		t.Errorf("array slicing - length: get %d, expect %d", actual, expected)
	}

	sender := []int{1, 2, 3, 4}
	var receiver [8]int
	copy(receiver[:4], sender)
	if actual, expected := receiver[:], []int{1, 2, 3, 4, 0, 0, 0, 0}; !slices.Equal(actual, expected) {
		t.Errorf("copy slice to sliced array:\nget\n%v\nexpect\n%v", actual, expected)
	}

}
