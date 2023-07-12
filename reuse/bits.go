package reuse

import (
	"fmt"
	"strconv"
)

// https://stackoverflow.com/a/13870865
func ShowBits(i int64, digits int) string {
	fmtString := "%0" + strconv.Itoa(digits) + "s"
	return fmt.Sprintf(fmtString, strconv.FormatInt(i, 2))
}

// https://stackoverflow.com/a/23192263
func SetBit(i int, pos uint) int {
	return i | (1 << pos)
}

func ClearBit(i int, pos uint) int {
	return i & ^(1 << pos)
}

func HasBit(i int, pos uint) bool {
	return i & (1 << pos) > 0
}