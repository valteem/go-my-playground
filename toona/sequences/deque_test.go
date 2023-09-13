package sequences_test

import (
	"fmt"
	"testing"

	"my.play.go/toona/sequences"
)

const (
	qCap = 5
)

func TestCircularDeque(t *testing.T) {

	q := sequences.NewCircularDeque[int](qCap)

	if !q.IsEmpty() {
		t.Error("wrong IsEmpty() result: should be true")
	}

	if _, e := q.DeleteFirst(); e == nil {
		t.Error("wrong DeleteFirst() result: should return error")
	}

	if _, e := q.DeleteLast(); e == nil {
		t.Error("wrong DeleteLast() result: should return error")
	}

	for i := 1; i <= qCap; i++ {
		q.AddFirst(i)
		q.AddLast(i)
	}

	fmt.Print(q)

	for i := qCap; i >= 1; i-- {
		u, _ := q.DeleteFirst()
		if u != i {
			t.Errorf("wrong DeleteFirst() result, get %v, should be %v", u, i)
		}
		v, _ := q.DeleteLast()
		if v != i {
			t.Errorf("wrong DeleteFirst() result, get %v, should be %v", v, i)
		}
	}

}