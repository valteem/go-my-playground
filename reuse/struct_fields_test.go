package reuse_test

import (
	"testing"
	"time"

	"github.com/valteem/reuse"
)

func TestStructFields(t *testing.T) {
	s := "some item"
	i := reuse.Item{s}
	time := time.Now()
	iw := reuse.ItemWrapper{i, time}
	if iw.Name != s {
		t.Errorf("Should be equal: %+v, %+v", iw.Name, s)
	}
	if iw.Time != time {
		t.Errorf("Should be equal: %+v, %+v", iw.Name, s)
	}
}