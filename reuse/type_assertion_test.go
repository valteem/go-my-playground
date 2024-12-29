package reuse_test

import (
	"reflect"
	"testing"

	"github.com/valteem/reuse"
)

func TestAssertType(t *testing.T) {

	tests := []struct {
		desc   string
		input  any
		output string
	}{
		{
			desc:   "Person",
			input:  reuse.Person{Name: "name", Age: 30},
			output: "Person",
		},
		{
			desc:   "StockItem",
			input:  reuse.StockItem{ID: 1, Description: "description", SupplierID: 1},
			output: "Stock Item",
		},
		{
			desc: "anonimous struct",
			input: struct {
				key   string
				value string
			}{
				key:   "somekey",
				value: "somevalue",
			},
			output: "Unknown",
		},
	}

	for _, tc := range tests {
		if output := reuse.AssertType(tc.input); output != tc.output {
			t.Errorf("%s: get %s, expect %s", tc.desc, output, tc.output)
		}
	}

	as := struct {
		key   string
		value string
	}{
		key:   "somekey",
		value: "somevalue",
	}
	typeOfAS := "struct { key string; value string }"
	if typeOf := reflect.TypeOf(as).String(); typeOf != typeOfAS {
		t.Errorf("reflect.Typeof(): get %s, expect %s", typeOf, typeOfAS)
	}

	//	fmt.Println(reflect.TypeOf(reflect.TypeOf(as))) // *reflect.rtype (type.go::317)

}

func TestInterfaceTypeAssertion(t *testing.T) {

	tests := []struct {
		input  any
		output string
	}{
		{reuse.TestObjectA{ID: 1}, "TypeObjectA"},
		{reuse.TestObjectB{ID: 1}, "TypeObjectB"},
		{"text", "Unknown"},
	}

	for _, tc := range tests {
		if output := reuse.TellObjectType(tc.input); output != tc.output {
			t.Errorf("get %s, expect %s", output, tc.output)
		}
	}

}

func TestCanSet(t *testing.T) {

	tests := []struct {
		input  any
		output bool
	}{
		{42, false},
		{"world", false},
		{true, false},
	}

	for _, tc := range tests {
		if output := reflect.ValueOf(tc.input).CanSet(); output != tc.output {
			t.Errorf("%v: get %t, expect %t", tc.input, output, tc.output)
		}
	}

}
