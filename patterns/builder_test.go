package patterns_test

import (
	"testing"

	"github.com/valteem/patterns"
)

func TestBuilder(t *testing.T) {
	s := patterns.NewScitation().Book("some book").Page(11).Text("some text")
	sg := s.Get()
	st := "some book 11 some text"
	if !(sg == st) {
		t.Error("wrong Get() result")
	}
}