package reuse_test

import (
	"fmt"
	"reflect"
	"testing"
)

type contract interface {
	showSet() string
}

type smaller struct {
	field1 int
	field2 int
}

func (s *smaller) showSet() string {
	return fmt.Sprintf("%d-%d", s.field1, s.field2)
}

type larger struct {
	field1 int
	field2 int
	// field5 int
}

func (l *larger) showSet() string {
	return fmt.Sprintf("%d-%d", l.field1, l.field2)
}
func TestTypeConversion(t *testing.T) {

	var c1 contract = (*larger)(nil)
	var l1 *larger = (*larger)(nil)
	// Cannot convert nil to struct value as nil only applies to:
	// pointer, channel, slice, map, interface
	// var s1 smaller = smaller(nil) // does not compile

	var s2 = &smaller{1, 2}
	var l2 *larger = (*larger)(s2)

	var s3 = smaller{1, 2}
	var l3 larger = larger(s3)

	// Detect nil pointer dereference (https://stackoverflow.com/a/50487104)

	if !(c1 == nil || reflect.ValueOf(c1).IsNil()) {
		if out := c1.showSet(); out != "0-0" {
			t.Errorf("interface type variable converted from nil: %s", out)
		}
	}

	if !(l1 == nil || reflect.ValueOf(c1).IsNil()) {
		if out := l1.showSet(); out != "0-0" {
			t.Errorf("struct variable converted from nil: %s", out)
		}
	}

	if out := l2.showSet(); out != "1-2" {
		t.Errorf("struct pointer variable converted from another struct pointer variable: %s", out)
	}

	if out := l3.showSet(); out != "1-2" {
		t.Errorf("struct value variable converted from another struct value variable: %s", out)
	}
}
