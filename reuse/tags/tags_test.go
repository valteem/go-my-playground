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

type reflectOutput struct {
	kind       reflect.Kind
	kindStr    string
	typeOfName string
	typeOfStr  string
	fieldNames []string
}

type exampleStruct struct {
	f1 string
	f2 int
}

func TestReflectBasic(t *testing.T) {

	tests := []struct {
		input  any
		output reflectOutput
	}{
		{
			input: struct {
				f1 string
				f2 int
			}{"field value", 1},
			output: reflectOutput{reflect.Struct, "struct", "", "struct { f1 string; f2 int }", []string{"f1", "f2"}},
		},
		{
			input:  exampleStruct{"field value", 1},
			output: reflectOutput{reflect.Struct, "struct", "exampleStruct", "tags.exampleStruct", []string{"f1", "f2"}},
		},
	}

	for _, tc := range tests {

		output := reflectOutput{}

		v := reflect.ValueOf(tc.input)
		output.kind = v.Kind()
		output.kindStr = v.Kind().String()

		output.typeOfName = reflect.TypeOf(tc.input).Name()
		output.typeOfStr = reflect.TypeOf(tc.input).String()

		for i := 0; i < v.NumField(); i++ {
			output.fieldNames = append(output.fieldNames, reflect.TypeOf(tc.input).Field(i).Name)
		}

		if !reflect.DeepEqual(output, tc.output) {
			t.Errorf("get\n%v\nexpect\n%v", output, tc.output)
		}

	}

}
