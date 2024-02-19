package reuse_test

import (
	"testing"

	iradix "github.com/hashicorp/go-immutable-radix/v2"
)

func TestBasicInsert(t *testing.T) {
	r := iradix.New[int]()
	r, _, _ = r.Insert([]byte("a"), 1)
	r, _, _ = r.Insert([]byte("ab"), 2)
	r, _, _ = r.Insert([]byte("ax"), 102)
	r, _, _ = r.Insert([]byte("abc"), 3)
	r, oldValue, isSet := r.Insert([]byte("abc"), 33)
	if oldValue != 3 {
		t.Errorf("Old value should be 11, is %+v", oldValue)
	}
	if r.Len() != 4{
		t.Errorf("Tree length should be 4, is %+v", r.Len())
	}
	if isSet != true {
		t.Errorf("Is Set flag should be true, is %+v", isSet)
	}

}