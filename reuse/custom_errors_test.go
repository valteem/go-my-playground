package reuse_test

import (
	"fmt"
	"testing"

	"github.com/valteem/reuse"
)

const (
	emsg = "some custom error"
	umsg = "unspecified error"
)

func ThrowCustomError() (string, error) {
	return "custom error", &reuse.CustomError{}
}
func TestCustomError(t *testing.T) {
	_, e := ThrowCustomError()
	if e != nil && e.Error() != "custom error" {
		t.Errorf("wrong error thrown: get %q, expect 'custom error'", e)
	}
}

func CastToCustomErrorPointer(err error) *reuse.AnotherCustomError {
	e, ok := err.(*reuse.AnotherCustomError)
	if !ok {
		return &reuse.AnotherCustomError{Err: fmt.Errorf("unspecified error")}
	}
	return e
}

func TestCastingCustomErrorPointer(t *testing.T) {

	e := fmt.Errorf(emsg)
	a := CastToCustomErrorPointer(e)
	if a.Err.Error() != umsg {
		t.Errorf("returns %q, expect %q", a.Err.Error(), umsg) // 'unspecified error'
	}

	c := &reuse.AnotherCustomError{Err: fmt.Errorf(emsg)} // reference to custom error !!!
	b := CastToCustomErrorPointer(c)
	if b.Err.Error() != emsg {
		t.Errorf("returns %q, expect %q", b.Err.Error(), emsg) // 'some custom error'
	}

}

func CastToCustomErrorValue(err error) reuse.AnotherCustomError {
	e, ok := err.(reuse.AnotherCustomError)
	if !ok {
		return reuse.AnotherCustomError{Err: fmt.Errorf("unspecified error")}
	}
	return e
}

func TestCastingCustomErrorValue(t *testing.T) {

	e := fmt.Errorf(emsg)
	a := CastToCustomErrorValue(e)
	if a.Err.Error() != umsg {
		t.Errorf("returns %q, expect %q", a.Err.Error(), umsg) // 'unspecified error'
	}

	c := reuse.AnotherCustomError{Err: fmt.Errorf(emsg)}
	b := CastToCustomErrorValue(c)
	if b.Err.Error() != emsg {
		t.Errorf("returns %q, expect %q", b.Err.Error(), emsg) // 'some custom error'
	}

}
