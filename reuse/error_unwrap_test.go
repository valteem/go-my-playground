package reuse_test

import (
	"errors"
	"fmt"
	"net/url"
	//	"reflect"
	"testing"

	"github.com/valteem/reuse"
)

func TestErrorUnwrap(t *testing.T) {
	err := errors.Unwrap(reuse.Outer(-1))
	fmt.Println(err)
	if errors.Is(err, reuse.ErrNotNegativeValue) {
		fmt.Println("sentinel error found")
	}
	if errors.Unwrap(err) == nil {
		fmt.Println("no more errors inside")
	}
}

func TestCustomErrorComparison(t *testing.T) {

	input := "some invalid URL string"

	e1 := &url.Error{
		Op:  "parse",
		URL: input,
		Err: errors.New("invalid URI for request")}

	e2 := &url.Error{
		Op:  "parse",
		URL: input,
		Err: errors.New("invalid URI for request")}

	if errors.Is(e1, e2) {
		t.Errorf("expect errors not considered equal")
	}

	// avoid using reflect.DeepEqual with errors
	/*
		if !reflect.DeepEqual(e1, e2) {
			t.Errorf("should be considered equal")
		}
	*/
}
