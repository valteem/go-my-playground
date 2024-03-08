package errors

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"strings"
)

type Code uint8

const (
	Unspecified Code = iota
	NotFound
	FailedPrecondition
	Internal
	AlreadyExists
	Unknown
)

func (c Code) String() string {
	switch c {
	case Unspecified:
		return "ERROR_CODE_UNSPECIFIED"
	case NotFound:
		return "NOT_FOUND"
	case FailedPrecondition:
		return "FAILED_PRECONDITION"
	case Internal:
		return "INTERNAL_ERROR"
	case AlreadyExists:
		return "ALREADY_EXISTS"
	case Unknown:
		return "UNKNOWN"
	}
	panic(fmt.Sprintf("unknown error code %d", c))
}

type Op string

type Error struct {
	Code Code
	Op   Op
	Err  error
}

// Return value include e.Op value
func (e *Error) DebugString() string {
	var b strings.Builder
	if e.Op != "" {
		b.WriteString(string(e.Op))
	}
	if e.Code != Unspecified {
		if b.Len() > 0 {
			b.WriteString(": ")
		}
		b.WriteString(e.Code.String())
	}
	if e.Err != nil {
		if b.Len() > 0 {
			b.WriteString(": ")
		}
		b.WriteString(e.Err.Error())
	}
	return b.String()
}

// implements general error interface, returns code and 'nested' error, skips ops
func (e *Error) Error() string {
	var b strings.Builder
	if e.Code != Unspecified {
		b.WriteString(e.Code.String())
	}
	if e.Err != nil {
		if b.Len() > 0 {
			b.WriteString(": ")
		}
		b.WriteString(e.Err.Error())
	}
	return b.String()
}

func E(args ...any) error {
	if len(args) == 0 {
		panic("calling error.E() with no arguments")
	}
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case Code:
			e.Code = arg
		case error:
			e.Err = arg
		case string:
			e.Err = errors.New(arg)
		default:
			_, file, line, _ := runtime.Caller(1)
			log.Printf("errors.E(): malformed call from %s:%d: %v", file, line, args)
			return fmt.Errorf("unknown type %T, value %v in error call", arg, arg)
		}
	}
	return e
}

func CanonicalCode(err error) Code {
	if err == nil {
		return Unspecified
	}
	e, ok := err.(*Error)
	if !ok {
		return Unspecified
	}
	if e.Code == Unspecified {
		return CanonicalCode(e.Err)
	}
	return e.Code
}
