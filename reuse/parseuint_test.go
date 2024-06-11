package reuse_test

import (
	"fmt"
	"strconv"
	"testing"
)

func customBaseReprConvertToDec(base int, input string) (uint64, error) {
	// Need to check that any digit value is not greater than (base - 1)
	l := len(input)
	if l == 0 {
		return 0, nil
	}
	digit := 1
	var output uint64 = 0
	for i := l - 1; i >= 0; i-- {
		dh, e := strconv.Atoi(string(input[i]))
		if e != nil {
			return 0, e
		}
		if dh >= base {
			return 0, fmt.Errorf("digit value should not exceed (base - 1)")
		}
		output += uint64(dh * digit)
		digit *= base
	}
	return output, nil
}

func TestCustomBaseReprToDec(t *testing.T) {
	tests := []struct {
		base   int
		repr   string
		output uint64
	}{
		{
			base:   16,
			repr:   "11",
			output: 17,
		},
		{
			base:   16,
			repr:   "111",
			output: 273, // 256 + 16 + 1
		},
		{
			base:   8,
			repr:   "111",
			output: 73, // 64 + 8 + 1
		},
		{
			base:   10,
			repr:   "1012141618",
			output: 1012141618,
		},
	}
	for _, tc := range tests {
		if output, _ := customBaseReprConvertToDec(tc.base, tc.repr); output != tc.output {
			t.Errorf("Get %d, expect %d", output, tc.output)
		}
	}
}
func TestParseUint(t *testing.T) {

	s := "1011"
	bases := []int{2, 4, 8, 10, 16, 32, 36} // base must not exceed 36

	for _, b := range bases {
		actualOutput, e := strconv.ParseUint(s, b, 64)
		if e != nil {
			t.Errorf("ParseUint() returns error %v for string %q and base %d", e, s, b)
		}
		expectedOutput, _ := customBaseReprConvertToDec(b, s)
		if actualOutput != expectedOutput {
			t.Errorf("Converting %q with base %d: get %d, expect %d", s, b, actualOutput, expectedOutput)
		}
	}

	// https://go.dev/ref/spec#Integer_literals
	var ordinary_int_repr uint64 = 12141618
	var literal_int_repr uint64 = 12_14_16_18
	if literal_int_repr != ordinary_int_repr {
		t.Errorf("%d %d", literal_int_repr, ordinary_int_repr)
	}

}
