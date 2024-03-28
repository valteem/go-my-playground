package reuse_test

import (
	"strconv"
	"testing"
)

func TestIntConversion(t *testing.T) {
	tests := []struct {
		description    string
		value          any
		formatFunc     func(any) string
		formattedValue string
	}{
		{
			description: "int8, decimal",
			value:       int8(16),
			formatFunc: func(value any) string {
				v := value.(int8)
				return strconv.FormatInt(int64(v), 10)
			},
			formattedValue: "16",
		},
		{
			description: "int8, octal",
			value:       int8(16),
			formatFunc: func(value any) string {
				v := value.(int8)
				return strconv.FormatInt(int64(v), 8)
			},
			formattedValue: "20",
		},
		{
			description: "int8, hexadecimal",
			value:       int8(16),
			formatFunc: func(value any) string {
				v := value.(int8)
				return strconv.FormatInt(int64(v), 16)
			},
			formattedValue: "10",
		},
		{
			description: "uint8, decimal",
			value:       int8(-16),
			formatFunc: func(value any) string {
				v := value.(int8)
				return strconv.FormatUint(uint64(v), 10)
			},
			formattedValue: strconv.FormatUint(^uint64(0)-16+1, 10),
		},
	}
	for _, tst := range tests {
		if v := tst.formatFunc(tst.value); v != tst.formattedValue {
			t.Errorf("%s: get %v, expect %v", tst.description, v, tst.formattedValue)
		}
	}
}

func TestIntToUint(t *testing.T) {
	tests := []struct {
		description string
		input       uint8
		//		convInput   func(uint8) int8
		output int8
	}{
		{
			input:  255,
			output: -1, // 255 - 128 = 127, -128 + 127 = -1
		},
	}
	for _, tst := range tests {
		if v := int8(tst.input); v != tst.output {
			t.Errorf("%s, get %d, expect %d", tst.description, v, tst.output)
		}
	}
}
