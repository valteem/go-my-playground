package testutil

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEquateInt64ApproxEqual(t *testing.T) {
	tests := []struct {
		description string
		v1          int64
		v2          int64
		margin      int64
	}{
		{
			description: "exact equality: zero margin",
			v1:          31,
			v2:          31,
			margin:      0,
		}, {
			description: "exact equality: non-zero margin",
			v1:          31,
			v2:          31,
			margin:      8,
		},
		{
			description: "equality within non-zero margin",
			v1:          31,
			v2:          35,
			margin:      8,
		},
	}

	for _, tst := range tests {
		opt := EquateInt64Approx(tst.margin)
		if !cmp.Equal(tst.v1, tst.v2, opt) {
			t.Errorf("%s: %v must be equal %v within margin %v", tst.description, tst.v1, tst.v2, tst.margin)
		}

	}
}

func TestEquateInt64ApproxNotEqual(t *testing.T) {
	tests := []struct {
		description string
		v1          int64
		v2          int64
		margin      int64
	}{
		{
			description: "exact equality: zero margin",
			v1:          31,
			v2:          32,
			margin:      0,
		}, {
			description: "exact equality: non-zero margin",
			v1:          31,
			v2:          33,
			margin:      1,
		},
		{
			description: "equality within non-zero margin",
			v1:          31,
			v2:          35,
			margin:      3,
		},
	}

	for _, tst := range tests {
		opt := EquateInt64Approx(tst.margin)
		if cmp.Equal(tst.v1, tst.v2, opt) {
			t.Errorf("%s: %v cannot be equal %v within margin %v", tst.description, tst.v1, tst.v2, tst.margin)
		}

	}
}
