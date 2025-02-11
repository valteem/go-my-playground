package encode

import (
	"encoding/binary"
	"math"
	"slices"
	"strconv"
	"testing"
)

func TestPutUint(t *testing.T) {

	tests := []struct {
		input    int
		outputBE []byte // most significant byte  - smallest memory address
		outputLE []byte // least significant byte - smallest memory address
	}{
		{
			input:    0b00000001_00000010_00000100_00001000,
			outputBE: []byte{1, 2, 4, 8},
			outputLE: []byte{8, 4, 2, 1},
		},
		{
			input:    0b10000000_01000000_00100000_00010000,
			outputBE: []byte{128, 64, 32, 16},
			outputLE: []byte{16, 32, 64, 128},
		},
	}

	for _, tc := range tests {

		input := uint32(tc.input)

		outputBE, outputLE := make([]byte, 4), make([]byte, 4)

		binary.BigEndian.PutUint32(outputBE, input)
		binary.LittleEndian.PutUint32(outputLE, input)

		if !slices.Equal(outputBE, tc.outputBE) {
			t.Errorf("Big Endian, input %b:\nget\n%v\nexpect\n%v\n", tc.input, outputBE, tc.outputBE)
		}

		if !slices.Equal(outputLE, tc.outputLE) {
			t.Errorf("Little Endian, input shift %d bits:\nget\n%v\nexpect\n%v\n", tc.input, outputLE, tc.outputLE)
		}

	}

}

func TestDecodeUInt(t *testing.T) {

	tests := []struct {
		input    []byte
		outputBE uint32
		outputLE uint32
	}{
		{
			input:    []byte{0, 0, 0, 1},
			outputBE: 1,
			outputLE: 1 << 24, // initial position 1 + shift 24 = position 25
		},
		{
			input:    []byte{0, 0, 1, 0},
			outputBE: 1 << 8,
			outputLE: 1 << 16,
		},
	}

	for _, tc := range tests {

		outputBE := binary.BigEndian.Uint32(tc.input)
		if outputBE != tc.outputBE {
			t.Errorf("Big Endian, decoding %v:\nget\n%d\nexpect\n%d\n", tc.input, outputBE, tc.outputBE)
		}

		outputLE := binary.LittleEndian.Uint32(tc.input)
		if outputLE != tc.outputLE {
			t.Errorf("Little Endian, decoding %v:\nget\n%d\nexpect\n%d\n", tc.input, outputLE, tc.outputLE)
		}
	}

}

func TestPutUvarint(t *testing.T) {

	tests := []struct {
		descr  string
		input  uint64
		output []byte
	}{
		{"8", 8, []byte{8, 0, 0, 0}}, // Little Endian
		{"256", 256, []byte{128, 2, 0, 0}},
		/*
				Step 0, input in binary format:
			        (00000001)(00000000)
				Step 1, divide input to 7-bit segments:
				    (0000010)(0000000)
				Step 2, reverse segment order:
					(0000000)(0000010)
				Step 4, add leading bit with 1 to every segment other the last one, bit with 0 to the last
				    (10000000)(00000010)
				Step 5 (representation only), convert output to sequence of bytes
				    (128)(2)
		*/
		{"257", 257, []byte{129, 2, 0, 0}},
		/*
			(00000001)(00000001)
			(0000010)(0000001)
			(0000001)(0000010)
			(10000001)(00000010)
			(129)(2)
		*/
		{"128*3", 128 * 3, []byte{128, 3, 0, 0}},
		{"128*3 + 64", 128*3 + 64, []byte{192, 3, 0, 0}},
		{"128 * 4 - 1", 128*4 - 1, []byte{255, 3, 0, 0}},
		{"512", 512, []byte{128, 4, 0, 0}}, // 128 * 4
		{"128 * 3", 128 * 3, []byte{128, 3, 0, 0}},
		{"128 * 128 * 2", 128 * 128 * 2, []byte{128, 128, 2, 0}},
		{"128 * 128 * 128 * 127", 128 * 128 * 128 * 127, []byte{128, 128, 128, 127}},
		{"128 * 128 * 128 * 128 - 1", 128*128*128*128 - 1, []byte{255, 255, 255, 127}},
		/*
				(00001111)(11111111)(11111111)(11111111)
				(1111111)(1111111)(1111111)(1111111)
				(1111111)(1111111)(1111111)(1111111)
				(11111111)(11111111)(11111111)(01111111)
			    (255)(255)(255)(127)
		*/
	}

	for _, tc := range tests {
		output := make([]byte, 4) // PutUvarint() panics if receiving buffer is too small
		binary.PutUvarint(output, tc.input)
		if !slices.Equal(output, tc.output) {
			t.Errorf("PutUvarint(%s):\nget\n%v\nexpect\n%v", tc.descr, output, tc.output)
		}
	}

}

func TestPutVarint(t *testing.T) {

	tests := []struct {
		descr  string
		input  int64
		output []byte
	}{
		{"8", 8, []byte{8, 0, 0, 0}},
		{"-8", -8, []byte{8, 0, 0, 0}},
	}

	for _, tc := range tests {
		output := make([]byte, 4)
		binary.PutVarint(output, tc.input)
		if !slices.Equal(output, tc.output) {
			t.Errorf("PutVarint(%s):\nget\n%v\nexpect\n%v\n", tc.descr, output, tc.output)
		}
	}

}

func TestBinaryXOR(t *testing.T) {

	tests := []struct {
		descr                      string
		input                      int64
		outputBinaryRepresentation string
		outputXOR                  int64
	}{
		{"8", 8, "1000", -9},
		{"-8", -8, "-1000", 7},
	}

	for _, tc := range tests {
		outputBinaryRepresentation := strconv.FormatInt(tc.input, 2)
		if outputBinaryRepresentation != tc.outputBinaryRepresentation {
			t.Errorf("binary representation of %q\nget\n%s\nexpect\n%s\n", tc.descr, outputBinaryRepresentation, tc.outputBinaryRepresentation)
		}
		outputXOR := ^tc.input
		if outputXOR != tc.outputXOR {
			t.Errorf("XOR of %q: get %d, expect %d", tc.descr, outputXOR, tc.outputXOR)
		}
	}
}

func TestBitwiseComplementInt64(t *testing.T) {

	tests := []struct {
		input int64
	}{
		{8},
		{-8},
	}

	for _, tc := range tests {
		outputUnary := ^tc.input
		outputBinary := (-1) ^ tc.input

		if outputBinary != outputUnary {
			t.Errorf("expect same bitwise XOR result for %d, get %d for binary and %d for unary",
				tc.input,
				outputBinary,
				outputUnary)
		}
	}

}

func TestBitwiseComplementUint64(t *testing.T) {

	tests := []struct {
		input  uint64
		output uint64
	}{
		{1, math.MaxUint64 - 1},
	}

	for _, tc := range tests {
		output := ^tc.input

		if output != tc.output {
			t.Errorf("expect same bitwise XOR result for %d, get %d expect %d",
				tc.input,
				output,
				tc.output)
		}
	}

}

func TestAllBitsTo1(t *testing.T) {

	tests := []struct {
		input  uint64
		output uint64
	}{
		{7, 7},
		{8, 15},
		{16, 31},
		{33, 63},
	}

	for _, tc := range tests {
		if output := AllSignificantBitsTo1(tc.input); output != tc.output {
			t.Errorf("AllBitsTo1(%d): get %d, expect %d", tc.input, output, tc.output)
		}
	}

}
