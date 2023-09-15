package sequences_test

import (
	"reflect"
	"testing"

	"my.play.go/toona/sequences"
)

func TestListDeque(t *testing.T) {

	q := sequences.NewListDeque[int]()

	if q.IsEmpty() != true {
		t.Error("wrong IsEmpty() result: should be true")
	}

	cap := 5
	for i := 1; i <= cap; i++ {
		q.AddFirst(i)
		q.AddLast(i)
	}
	for i := cap; i >= 1; i-- {
		if v, e := q.DeleteFirst(); v != i || e != nil {
			t.Errorf("wrong DeleteFirst() result: should be %v %v", v, e)
		}
		if v, e := q.DeleteLast(); v != i || e != nil {
			t.Errorf("wrong DeleteLast() result: should be %v %v", v, e)
		}		
	}

	v, e := q.DeleteFirst()
	if v != 0 || !reflect.DeepEqual(e.Error(), "empty deque") {
		t.Errorf("wrong DeleteFirst() result on empty deque")
	}
	v, e = q.DeleteLast()
	if v != 0 || !reflect.DeepEqual(e.Error(), "empty deque") {
		t.Errorf("wrong DeleteLast() result on empty deque")
	}
}