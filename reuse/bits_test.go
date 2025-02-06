package reuse_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/valteem/reuse"
)

func TestShowBits(t *testing.T) {
	for i := 0; i <= 8; i++ {
		fmt.Println(reuse.ShowBits(int64(i), 4))
	}
}

func TestSetBit(t *testing.T) {
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 4; j++ {
			fmt.Println(i, j, fmt.Sprintf("%05s", strconv.FormatInt(int64(reuse.SetBit(i, uint(j))), 2)))
		}
		fmt.Println(strings.Repeat("-", 10))
	}
}

func TestClearBit(t *testing.T) {
	i := (1 << 8) - 1
	for j := 0; j < 8; j++ {
		fmt.Println(i, j, reuse.ClearBit(i, uint(j)), fmt.Sprintf("%08s", strconv.FormatInt(int64(reuse.ClearBit(i, uint(j))), 2)))
	}
}

func TestHasBit(t *testing.T) {
	i, _ := strconv.ParseInt("10101010", 2, 64)
	for j := 0; j < 8; j++ {
		fmt.Println(i, fmt.Sprintf("%08s", strconv.FormatInt(int64(i), 2)), j, reuse.HasBit(int(i), uint(j)))
	}
}

func TestBitShiftRight(t *testing.T) {

	tests := []struct {
		input  int
		shift  int
		output int
	}{
		{1 << 32, 32, 1},
		{1 << 32, 33, 0},
		{1 << 32, 16, 1 << 16},
	}

	for _, tc := range tests {
		output := tc.input >> tc.shift
		if output != tc.output {
			t.Errorf("%d >> %d: get %d, expect %d", tc.input, tc.shift, output, tc.output)
		}
	}
}
