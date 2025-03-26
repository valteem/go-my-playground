package errortypeassert

import (
	"strconv"
)

var (
	ErrStr = NewStrError("string error")
	ErrInt = NewIntError(42)
)

type StrError struct {
	s string
}

func (e *StrError) Error() string {
	return e.s
}

func NewStrError(s string) *StrError {
	return &StrError{s: s}
}

type IntError struct {
	i int
}

func (e *IntError) Error() string {
	return strconv.Itoa(e.i)
}

func NewIntError(i int) *IntError {
	return &IntError{i: i}
}

type CompositeError struct {
	s string
	i int
}

func (e *CompositeError) Error() string {
	return e.s + "/" + strconv.Itoa(e.i)
}

func NewCompositeError(s string, i int) *CompositeError {
	return &CompositeError{s: s, i: i}
}

type AssertErrorType func(e error) (error, bool)
