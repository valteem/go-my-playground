// https://go.dev/blog/go1.13-errors
package reuse

import (
	"errors"
	"fmt"
)

const ErrNotNegativeValueMessage = "negative values not accepted" 

var ErrNotNegativeValue = errors.New(ErrNotNegativeValueMessage)

func inner(i int) error {
	if i < 0 {
		return ErrNotNegativeValue
	} else {
		return nil
	}
}

func Outer(i int) error {
	return fmt.Errorf("%v : %w", ErrNotNegativeValueMessage, inner(i))
}

