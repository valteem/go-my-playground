package reuse_test

import (
	"testing"

	"github.com/valteem/reuse"
)

func TestConvertUnixTiime(t *testing.T) {
	t1, t2 := reuse.ConvertUnixTimeZero()
	if t1 != t2 {
		t.Errorf("%+v %+v", t1, t2)
	}
}