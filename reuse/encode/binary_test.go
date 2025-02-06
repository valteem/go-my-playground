package encode

import (
	"encoding/binary"
	"slices"
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
