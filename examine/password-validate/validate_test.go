package main

import (
	"testing"

	"github.com/go-passwd/validator"
)

func TestValidate(t *testing.T) {

	tests := []struct {
		input     string
		validator validator.Validator
		isErr     bool
		outputMsg string
	}{
		{
			input:     "password",
			validator: validator.Validator{validator.ContainsOnly("adoprsw", nil)},
			isErr:     false,
		},
		{
			input:     "password",
			validator: validator.Validator{validator.ContainsOnly("adoprs", nil)},
			isErr:     true,
		},
		{
			input: "password",
			validator: validator.Validator{
				validator.ContainsOnly("adoprsw", nil),
				validator.StartsWith("a", nil),
			},
			isErr: true,
		},
		{
			input: "password",
			validator: validator.Validator{
				validator.ContainsOnly("adoprsw", nil),
				validator.StartsWith("p", nil),
				validator.MinLength(10, nil),
			},
			isErr: true,
		},
	}

	var isErr bool
	for _, tc := range tests {

		err := tc.validator.Validate(tc.input)

		if err != nil {
			isErr = true
		} else {
			isErr = false
		}

		if isErr != tc.isErr {
			t.Errorf("validating %q with %v: get %t, expect %t", tc.input, tc.validator, isErr, tc.isErr)
		}
	}

}
