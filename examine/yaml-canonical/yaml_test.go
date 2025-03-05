package main

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestBasic(t *testing.T) {

	tests := []struct {
		input  string
		output *Person
	}{
		{
			input: `
name: Some Name
age: 42
address:
  city: Some City
  street: Some Street
  zipcode: 12345
`,
			output: &Person{
				Name: "Some Name",
				Age:  42,
				Address: &Address{
					City:    "Some City",
					Street:  "Some Street",
					ZipCode: "12345",
				},
			},
		},
	}

	for _, tc := range tests {

		output := &Person{}

		err := yaml.Unmarshal([]byte(tc.input), output)
		if err != nil {
			t.Fatalf("failed to unmarshal input YAML %q:\n %v", tc.input, err)
		}

		if !reflect.DeepEqual(output, tc.output) {
			t.Errorf("unmarhsling %q YAML:\nget\n%v\nexpect\n%v\n", tc.input, output, tc.output)
		}
	}

}

func TestEmbedded(t *testing.T) {

	tests := []struct {
		input  string
		output *PersonEmbeddedAddress
	}{
		{
			input: `
name: Some Name
age: 42
city: Some City
street: Some Street
zipcode: 12345		
`,
			output: &PersonEmbeddedAddress{
				Name: "Some Name",
				Age:  42,
				Address: Address{
					City:    "Some City",
					Street:  "Some Street",
					ZipCode: "12345",
				},
			},
		},
	}

	for _, tc := range tests {

		output := &PersonEmbeddedAddress{}

		err := yaml.Unmarshal([]byte(tc.input), output)
		if err != nil {
			t.Fatalf("failed to unmarshal input YAML %q:\n%v", tc.input, err)
		}

		if !reflect.DeepEqual(output, tc.output) {
			t.Errorf("unmarshaling %q YAML:\nget\n%v\nexpect\n%v\n", tc.input, output, tc.output)
		}
	}

}
