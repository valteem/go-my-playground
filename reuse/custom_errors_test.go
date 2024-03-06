package reuse_test

import (
	"testing"

	"github.com/valteem/reuse"
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
