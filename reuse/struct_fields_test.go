package reuse_test

import (
	"reflect"
	"slices"
	"testing"
	"time"

	"github.com/valteem/reuse"
)

func TestStructFields(t *testing.T) {
	s := "some item"
	i := reuse.Item{s}
	time := time.Now()
	iw := reuse.ItemWrapper{i, time}
	if iw.Name != s {
		t.Errorf("Should be equal: %+v, %+v", iw.Name, s)
	}
	if iw.Time != time {
		t.Errorf("Should be equal: %+v, %+v", iw.Name, s)
	}
}

type fieldOffset struct {
	name   string
	offset uintptr
}

/*
https://stackoverflow.com/a/65878722

Strings in Go are represented by reflect.StringHeader containing a pointer to actual string data
and a length of string:

type StringHeader struct {
	Data uintptr
	Len  int
}

unsafe.Sizeof(s) will only return the size of StringHeader struct but not the pointed data itself.
So it will be sum of 8 bytes for Data and 8 bytes for Len making it 16 bytes
*/

func TestStructFieldOffset(t *testing.T) {

	tests := []struct {
		input  any
		output []fieldOffset
	}{
		{
			input: struct {
				f1 string
				f2 int
			}{
				f1: "very-vey-very-long-field",
				f2: 2,
			},
			output: []fieldOffset{
				{"f1", 0}, {"f2", 16},
			},
		},
		{
			input: struct {
				f0 int64
				f1 string
				f2 int
			}{
				f0: 42,
				f1: "some text of arbitrary size",
				f2: 42,
			},
			output: []fieldOffset{
				{"f0", 0}, {"f1", 8}, {"f2", 24},
			},
		},
	}

	for _, tc := range tests {

		output := []fieldOffset{}

		v := reflect.ValueOf(tc.input)
		for i := 0; i < v.NumField(); i++ {
			field := fieldOffset{}
			field.name = v.Type().Field(i).Name
			field.offset = v.Type().Field(i).Offset
			output = append(output, field)
		}

		if !slices.Equal(output, tc.output) {
			t.Errorf("%v:\nget\n%v\nexpect\n%v\n", tc.input, output, tc.output)
		}
	}

}
