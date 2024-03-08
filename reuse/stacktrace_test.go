package reuse_test

import (
	"testing"

	"github.com/valteem/reuse"
)

func _A() string {
	return B()
}

func B() string {
	return C()
}

func C() string {
	return reuse.StackTrace()
}

func TestStackTrace(t *testing.T) {
	s := _A()
	if s == "" {
		t.Errorf("stack trace should not be empty")
	}

}
