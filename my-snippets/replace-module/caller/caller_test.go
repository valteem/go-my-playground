package caller

import (
	"testing"

	m "github.com/some-author/some-module"
)

func TestModuleReplace(t *testing.T) {

	tests := []struct {
		description string
		inputString string
		inputNumber int
		output      string
	}{
		{
			description: "negative input",
			inputString: "text",
			inputNumber: -1,
			output:      m.InputNegative,
		},
		{
			description: "one-digit input",
			inputString: "text",
			inputNumber: 1,
			output:      "text01",
		},
		{
			description: "two-digit input",
			inputString: "text",
			inputNumber: 11,
			output:      "text11",
		},
		{
			description: "large input",
			inputString: "text",
			inputNumber: 111,
			output:      m.InputTooLarge,
		},
	}

	for _, tc := range tests {
		output := m.Concat(tc.inputString, tc.inputNumber)
		if output != tc.output {
			t.Errorf("%s: get %q, expect %q", tc.description, output, tc.output)
		}
	}

}
