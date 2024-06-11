package model

import (
	"sort"
	"testing"
)

func TestFingerprintFromString(t *testing.T) {
	tests := []struct {
		input  string
		output Fingerprint
	}{
		{
			input:  "1234567890",
			output: 78187493520,
		},
		{
			input:  "1234567890abc",
			output: 320255973460668,
		},
	}
	for _, tc := range tests {
		output, err := FingerprintFromString(tc.input)
		if err != nil {
			t.Errorf("FingerprintFromString(%q): %v", tc.input, err)
		}
		if output != tc.output {
			t.Errorf("FingerprintFromString(%q): get %d, expect %d", tc.input, output, tc.output)
		}
	}
	for _, tc := range tests {
		output, err := ParseFingerprint(tc.input)
		if err != nil {
			t.Errorf("ParseFingerprint(%q): %v", tc.input, err)
		}
		if output != tc.output {
			t.Errorf("ParseFingerprint(%q): get %d, expect %d", tc.input, output, tc.output)
		}
	}
}

func TestFingerprintToString(t *testing.T) {
	tests := []struct {
		input  Fingerprint
		output string
	}{
		{
			input:  1234567890,
			output: "00000000499602d2",
		},
		{
			input:  2147483647,
			output: "000000007fffffff",
		},
		{
			input:  18_446_744_073_709_551_615, // integer literal
			output: "ffffffffffffffff",
		},
	}
	for _, tc := range tests {
		if output := tc.input.String(); output != tc.output {
			t.Errorf("String(%d): get %s, expect %s", tc.input, output, tc.output)
		}
	}
}

func TestFingerprintsSort(t *testing.T) {
	input := Fingerprints{
		12345678,
		1234,
		12,
		1234567,
	}
	output := Fingerprints{
		12,
		1234,
		1234567,
		12345678,
	}
	sort.Sort(input)
	for i, v := range input {
		if v != output[i] {
			t.Errorf("Sort(): get input[%d] = %d, expect %d", i, v, output[i])
		}
	}
}

func TestFingerprintSetEqual(t *testing.T) {
	tests := []struct {
		descr  string
		fs1    FingerprintSet
		fs2    FingerprintSet
		output bool
	}{
		{
			descr: "identical fingerprint sets",
			fs1: FingerprintSet{
				12345:    struct{}{},
				1234567:  struct{}{},
				12345678: struct{}{},
			},
			fs2: FingerprintSet{
				12345:    struct{}{},
				1234567:  struct{}{},
				12345678: struct{}{},
			},
			output: true,
		},
		{
			descr: "same length. different keys",
			fs1: FingerprintSet{
				12345:    struct{}{},
				1234567:  struct{}{},
				12345678: struct{}{},
			},
			fs2: FingerprintSet{
				123:      struct{}{},
				1234567:  struct{}{},
				12345678: struct{}{},
			},
			output: false,
		},
		{
			descr: "different length",
			fs1: FingerprintSet{
				12345:    struct{}{},
				1234567:  struct{}{},
				12345678: struct{}{},
			},
			fs2: FingerprintSet{
				1234567:  struct{}{},
				12345678: struct{}{},
			},
			output: false,
		},
	}
	for _, tc := range tests {
		if output := tc.fs1.Equal(tc.fs2); output != tc.output {
			t.Errorf("%s: Equal() should return %t, get %t", tc.descr, tc.output, output)
		}
	}
}

func TestFingerprintSetIntersection(t *testing.T) {
	tests := []struct {
		descr  string
		input1 FingerprintSet
		input2 FingerprintSet
		output FingerprintSet
	}{
		{
			descr:  "two empty sets",
			input1: FingerprintSet{},
			input2: FingerprintSet{},
			output: FingerprintSet{},
		},
		{
			descr:  "one empty set",
			input1: FingerprintSet{},
			input2: FingerprintSet{
				1:   struct{}{},
				12:  struct{}{},
				123: struct{}{},
			},
			output: FingerprintSet{},
		},
		{
			descr: "equal sets",
			input1: FingerprintSet{
				1:   struct{}{},
				12:  struct{}{},
				123: struct{}{},
			},
			input2: FingerprintSet{
				1:   struct{}{},
				12:  struct{}{},
				123: struct{}{},
			},
			output: FingerprintSet{
				1:   struct{}{},
				12:  struct{}{},
				123: struct{}{},
			},
		},
		{
			descr: "non-empty non-equal sets, same length",
			input1: FingerprintSet{
				1:   struct{}{},
				12:  struct{}{},
				123: struct{}{},
			},
			input2: FingerprintSet{
				1:    struct{}{},
				12:   struct{}{},
				1234: struct{}{},
			},
			output: FingerprintSet{
				1:  struct{}{},
				12: struct{}{},
			},
		},
		{
			descr: "non-empty non-equal sets, different length",
			input1: FingerprintSet{
				1:   struct{}{},
				12:  struct{}{},
				123: struct{}{},
			},
			input2: FingerprintSet{
				1:    struct{}{},
				12:   struct{}{},
				123:  struct{}{},
				1234: struct{}{},
			},
			output: FingerprintSet{
				1:   struct{}{},
				12:  struct{}{},
				123: struct{}{},
			},
		},
	}
	for _, tc := range tests {
		if output := tc.input1.Intersection(tc.input2); !output.Equal(tc.output) {
			t.Errorf("%s: Intersection() returns %v, expect %v", tc.descr, output, tc.output)
		}
	}
}
