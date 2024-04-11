package reuse_test

import (
	"testing"

	list "github.com/emirpasic/gods/v2/lists/arraylist"
)

func TestIteratorBeginEndFirstLastSingleElement(t *testing.T) {
	list := list.New[string]("a")
	it := list.Iterator()

	idxBegin, idxFirst, idxLast, idxEnd := -1, 0, 0, 1

	it.Begin()
	if i := it.Index(); i != idxBegin {
		t.Errorf("it.Index() on single element list should return %d after it.Begin(), get %d", idxBegin, i)
	}

	it.First()
	if i := it.Index(); i != idxFirst {
		t.Errorf("it.Index() on single element list should return %d after it.First(), get %d", idxFirst, i)
	}

	it.Last()
	if i := it.Index(); i != idxLast {
		t.Errorf("it.Index() on single element list should return %d after it.Last(), get %d", idxLast, i)
	}

	it.End()
	if i := it.Index(); i != idxEnd {
		t.Errorf("it.Index() on single element list should return %d after it.End(), get %d", idxEnd, i)
	}
}