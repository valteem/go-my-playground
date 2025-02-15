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
