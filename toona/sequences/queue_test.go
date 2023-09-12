package sequences_test

import (
	"fmt"
	"testing"

	"my.play.go/toona/sequences"
)

const (
	queueCap = 5
)

func TestCircularQueue(t *testing.T) {

	q := sequences.NewCircularQueue[int](queueCap)

	if q.IsEmpty() == false {
		t.Error("wrong isEmpty result: should be true")
	}

	for i := 1; i <= queueCap; i++ {
		q.Enqueue(i)
	}

	if q.IsFull() == false {
		t.Error("wrong isFull result: should be true")
	}

	fmt.Println(q)

	for i := 1; i <= queueCap; i++ {
		_ = q.Dequeue()
	}

	if q.IsEmpty() == false {
		t.Error("wrong isEmpty result: should be true")
	}

	fmt.Println(q)

}

func TestCircularQueueResize(t *testing.T) {

	q := sequences.NewCircularQueue[int](queueCap)

	for i := 1; i <= queueCap * 2; i++ {
		q.Enqueue(i)
	}

	for i := 1; i <= queueCap * 2; i++ {
		if v := q.Dequeue(); v != i {
			t.Errorf("wrong Dequeue() result: got %v expected %v", v, i)
		}
	}

}
