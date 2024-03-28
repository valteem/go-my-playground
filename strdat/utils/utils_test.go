package utils

import (
	"testing"
)

func TestToString(t *testing.T) {
	tests := []struct {
		description string
		input       any
		output      string
	}{
		{
			description: "int8",
			input:       int8(127),
			output:      "127",
		},
		{
			description: "int16",
			input:       int16(1<<15 - 1),
			output:      "32767",
		},
		{
			description: "int32",
			input:       int32(1<<31 - 1),
			output:      "2147483647",
		},
		{
			description: "int64",
			input:       int64(1<<63 - 1),
			output:      "9223372036854775807",
		},
		{
			description: "uint8",
			input:       uint8(255),
			output:      "255",
		},
		{
			description: "uint16",
			input:       uint16(1<<16 - 1),
			output:      "65535",
		},
		{
			description: "uint32",
			input:       uint32(1<<32 - 1),
			output:      "4294967295",
		},
		{
			description: "uint64",
			input:       uint64(1<<64 - 1),
			output:      "18446744073709551615",
		},
		{
			description: "string",
			input:       "custom string",
			output:      "custom string",
		},
		{
			description: "bool true",
			input:       true,
			output:      "true",
		},
		{
			description: "bool false",
			input:       false,
			output:      "false",
		},
	}
	for _, tst := range tests {
		if v := ToString(tst.input); v != tst.output {
			t.Errorf("%s, get %v, expect %s", tst.description, v, tst.output)
		}
	}

}
