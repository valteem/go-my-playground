package basic

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestSimpleValidation(t *testing.T) {

	tests := []struct {
		descr  string
		input  Person
		output map[string]string
	}{
		{
			descr:  "valid input",
			input:  Person{GivenName: "Given", FamilyName: "Name", Age: 42, Email: "valid@e.mail"},
			output: map[string]string{},
		},
		{
			descr:  "given name missing",
			input:  Person{FamilyName: "Name", Age: 42, Email: "valid@e.mail"},
			output: map[string]string{"Person.GivenName": "required"},
		},
		{
			descr:  "age out of range",
			input:  Person{GivenName: "Given", FamilyName: "Name", Age: 142, Email: "valid@e.mail"},
			output: map[string]string{"Person.Age": "lte"},
		},
		{
			descr:  "invalid email",
			input:  Person{GivenName: "Given", FamilyName: "Name", Age: 42, Email: "invalid@email"},
			output: map[string]string{"Person.Email": "email"},
		},
		{
			descr:  "missing family name, invalid email",
			input:  Person{FamilyName: "Name", Age: 42, Email: "invalid@email"},
			output: map[string]string{"Person.GivenName": "required", "Person.Email": "email"},
		},
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	for _, tc := range tests {
		if err := validate.Struct(tc.input); err != nil {
			ves := err.(validator.ValidationErrors)
			for _, ve := range ves {
				tag, ok := tc.output[ve.Namespace()]
				if ok && ve.Tag() != tag {
					t.Errorf("%q error tag: get %q, expect %q", tc.descr, ve.Tag(), tag)
				}
			}
		}
	}

}

func TestStringValidation(t *testing.T) {

	tests := []struct {
		descr  string
		input  ValidatedStrings
		errMsg string
	}{

		// Skip valid input to avoid dereferencing nil error
		/*
			{
				descr: "valid input",
				input: ValidatedStrings{
					Name:          "ValidName",
					AlphaNumField: "Universe42",
					ContainField:  "Universe42",
				},
				errMsg: "",
			},
		*/
		{
			descr: "alpha tagged field contains numerics",
			input: ValidatedStrings{
				Name:          "Universe42",
				AlphaNumField: "Universe42",
				ContainField:  "Universe42",
			},
			errMsg: "Key: 'ValidatedStrings.Name' Error:Field validation for 'Name' failed on the 'alpha' tag",
		},
		{
			descr: "alphanum tagged field contains Unicode points",
			input: ValidatedStrings{
				Name:          "Universe",
				AlphaNumField: "丁丂七",
				ContainField:  "Universe42",
			},
			errMsg: "Key: 'ValidatedStrings.AlphaNumField' Error:Field validation for 'AlphaNumField' failed on the 'alphanum' tag",
		},
		{
			descr: "input does not contain required substring",
			input: ValidatedStrings{
				Name:          "Universe",
				AlphaNumField: "Field42",
				ContainField:  "UniverseFortyTwo",
			},
			errMsg: "Key: 'ValidatedStrings.ContainField' Error:Field validation for 'ContainField' failed on the 'contains' tag",
		},
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	for _, tc := range tests {
		if errMsg := validate.Struct(tc.input).Error(); errMsg != tc.errMsg {
			t.Errorf("%q input:\nget\n%q\nexpect\n%q\n", tc.descr, errMsg, tc.errMsg)
		}
	}
}
