// https://victoriametrics.com/blog/go-array/
package reuse_test

import (
	"reflect"
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
