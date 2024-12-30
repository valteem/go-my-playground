package main

import (
	"reflect"
	"testing"
)

func TestHeadTail(t *testing.T) {

	tests := []struct {
		input string
		sep   string
		head  string
		tail  string
	}{
		{"more sweet apples", "app", "more sweet ", "les"},
		{"know your customer", "x", "know your customer", ""},
	}

	for _, tc := range tests {
		if head, tail := HeadTail(tc.input, tc.sep); head != tc.head || tail != tc.tail {
			t.Errorf("head/tail for %s:%s: get  %s/%s, expect %s/%s", tc.input, tc.sep, head, tail, tc.head, tc.tail)
		}
	}
}

// https://stackoverflow.com/a/48874650
// Interface type variables are not addressable and cannot be set with Value.Set() methods
func TestSetWithProperType(t *testing.T) {

	field := reflect.StructField{}

	targetStr := ""
	inputStr := "apples"
	err := SetWithProperType(inputStr, reflect.ValueOf(&targetStr).Elem(), field)
	if err != nil {
		t.Fatalf("failed to assign %q: %v", inputStr, err)
	}
	if targetStr != inputStr {
		t.Errorf("get %q, expect %q", targetStr, inputStr)
	}

	targetInt := 0
	inputInt := "42"
	err = SetWithProperType(inputInt, reflect.ValueOf(&targetInt).Elem(), field)
	if err != nil {
		t.Fatalf("failed to assign %q: %v", inputStr, err)
	}
	if targetInt != 42 {
		t.Errorf("get %d, expect %d", targetInt, 42)
	}
}
