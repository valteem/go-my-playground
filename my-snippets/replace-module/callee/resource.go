package resource

import (
	"fmt"
)

// return values for non-accepted inputs
const (
	InputNegative = "input number cannot be negative"
	InputTooLarge = "input number too large"
)

func Concat(s string, i int) string {
	if i < 0 {
		return InputNegative
	}
	if i > 99 {
		return InputTooLarge
	}
	return s + fmt.Sprintf("%02d", i)
}
