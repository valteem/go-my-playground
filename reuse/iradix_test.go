package reuse_test

import (
	"testing"

	iradix "github.com/hashicorp/go-immutable-radix/v2"
)

func TestBasicInsert(t *testing.T) {
	r := iradix.New[int]()
	r, _, _ = r.Insert([]byte("abc"), 11)
	r, oldValue, isSet := r.Insert([]byte("abc"), 12)
	if oldValue != 11 {
		t.Errorf("Old value should be 11, is %+v", oldValue)
	}
	if r.Len() != 1{
		t.Errorf("Tree length should be 1, is %+v", r.Len())
	}
	if isSet != true {
		t.Errorf("Is Set flag should be true, is %+v", isSet)
	}

}