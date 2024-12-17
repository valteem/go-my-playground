package tags

import (
	"reflect"
	"slices"
	"testing"
)

func TestTags(t *testing.T) {

	tests := []struct {
		input  any
		tag    string
		output []string
	}{
		{
			input: struct {
				f1 string `customtag:"f1"`
				f2 int    `customtag:"f2"`
			}{
				f1: "f1value",
				f2: 1,
			},
			tag:    "customtag",
			output: []string{"f1", "f2"},
		},
		{
			input: struct {
				f1 string `customtag1:"f1"`
				f2 int    `customtag2:"f2"`
				f3 bool   `customtag1:"f3"`
			}{
				f1: "f1value",
				f2: 1,
				f3: true,
			},
			tag:    "customtag1",
			output: []string{"f1", "", "f3"},
		},
	}

	for _, tc := range tests {
		output := make([]string, 0)
		v := reflect.ValueOf(tc.input)
		for i := 0; i < v.NumField(); i++ {
			output = append(output, string(v.Type().Field(i).Tag.Get(tc.tag)))
		}
		if !slices.Equal(output, tc.output) {
			t.Errorf("%v tags:\nget\n%v\nexpect\n%v", tc.input, output, tc.output)
		}
	}
}
