package sequences_test

import (
	"reflect"
	"testing"

	"my.play.go/toona/sequences"
)

func TestListQueue(t *testing.T) {

	q := sequences.NewListQueue[int]()

	if q.IsEmpty() != true {
		t.Error("wrong IsEmpty() result: should be true")
	}

	cap := 10

	for i := 1; i <= cap; i++ {
		q.Enqueue(i)
	}
	for i := 1; i <= cap; i++ {
		v, e := q.Dequeue()
		if v != i || e != nil {
			t.Error("wrong Dequeue() result")
		}
	}

	v, e := q.Dequeue()
	if v != 0 || !reflect.DeepEqual(e.Error(), "empty queue") {
		t.Error("wrong Dequeue() result on empty queue")
	}
}