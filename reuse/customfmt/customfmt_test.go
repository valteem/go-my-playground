package customfmt

import (
	"fmt"
	"testing"
)

func TestGoString(t *testing.T) {

	tests := []struct {
		input  Building
		output string
	}{
		{
			input: Building{
				TotalArea:      1000,
				NumberOfFloors: 4,
				Addr: &Address{
					ZipCode: "12345",
					City:    "Dubuque, Iowa",
					Street:  "2194 Bennett St",
				},
			},
			output: "Total area 1000.0 sq.m., number of floors 4, Address: zip code 12345, city Dubuque, Iowa, street 2194 Bennett St",
		},
	}

	for _, tc := range tests {
		// The GoString method is used to print values passed as an operand to a %#v format.
		if output := fmt.Sprintf("%#v", tc.input); output != tc.output {
			t.Errorf("get\n%s\nexpect\n%s\n", output, tc.output)
		}
	}

}
