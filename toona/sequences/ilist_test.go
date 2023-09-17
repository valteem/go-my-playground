package sequences_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"my.play.go/toona/sequences"
)

func TestIList(t *testing.T) {

	l := sequences.NewIList[int]()

	if l.IsEmpty() != true {
		t.Error("wrong IsEmpty() result: should be true")
	}

	v := 0
	n := l.Insert(v, l.Tail(), l.Head())
	if e := n.Element; e != 0 {
		t.Errorf("wrong Insert result: expect %v, get %v\n", v, e)
	}

	cap := 5
	n = l.Tail().Next
	for i := 1; i <= cap; i++ {
		n = l.Insert(i, n, n.Next)
	}
	s := strings.TrimSuffix(fmt.Sprintln(l), "\n")
	if !reflect.DeepEqual(s, "tail<->0<->1<->2<->3<->4<->5<->head") {
		t.Error("wrong String() result")
	}

}