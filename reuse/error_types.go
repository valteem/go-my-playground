// Defines several error types

package reuse

import (
	"fmt"
)

const SignErrorMsg = "value cannot be negative"

type SignError struct{}

func (e *SignError) Error() string {
	return SignErrorMsg
}

type RangeError struct {
	LowerBound int
	UpperBound int
}

func (e *RangeError) Error() string {
	return fmt.Sprintf("value should be between %d and %d", e.LowerBound, e.UpperBound)
}
