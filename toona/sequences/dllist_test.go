package sequences_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"my.play.go/toona/sequences"
)

func TestDoubleLinkedList(t *testing.T) {

	l := sequences.NewDoubleLinkedList[int]()

	if _, e := l.RemoveFromHead(); e == nil {
		t.Error("wrong RemoveFromHead() result: should not be nil")
	}
	
	if _, e := l.RemoveFromTail(); e == nil {
		t.Error("wrong RemoveFromTail() result: should not be nil")
	}

	for i := 1; i <= 4; i++ {
		l.AddToHead(i)
		l.AddToTail(i)
	}

	s := strings.TrimSuffix(fmt.Sprintln(l), "\n")
	if !reflect.DeepEqual(s, "head<->4<->3<->2<->1<->1<->2<->3<->4<->tail") {
		t.Errorf("wrong String() result: %s", s)
	}

	for i := 4; i >= 1; i-- {
		if v, e := l.RemoveFromHead(); v != i || e != nil {
			t.Errorf("wrong RemoveFromHead() result: %+v, %+v", v, e)
		}
		if v, e := l.RemoveFromTail(); v != i || e != nil {
			t.Errorf("wrong RemoveFromTail() result: %+v, %+v", v, e)
		}
	}

	if l.IsEmpty() == false {
		t.Error("wrong IsEmpty() result: should be true")
	}
	
}