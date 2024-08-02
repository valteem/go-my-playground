package caller

import (
	"testing"

	"my.snippets/nce/resource"
)

func TestNestedExport(t *testing.T) {

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
			output:      resource.InputNegative,
		},
		{
			description: "one-digit input",
			inputString: "text",
			inputNumber: 1,
			output:      "text01",
		},
	}

	for _, tc := range tests {
		output := resource.Concat(tc.inputString, tc.inputNumber)
		if output != tc.output {
			t.Errorf("%s: get %s, expect %s", tc.description, output, tc.output)
		}
	}

}
