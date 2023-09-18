package sequences_test

import (
	"reflect"
	"testing"

	"my.play.go/toona/sequences"
)

func TestBasic(t *testing.T) {

	l := sequences.NewPList[int]()
	p := l.InsertToHead(1)
	if e := p.Element(); e != 1 {
		t.Errorf("wrong InsertTohead result: expect %v, get %v", 1, e)
	}
}

func TestAfter(t *testing.T) {
	l := sequences.NewPList[string]()
	v := []string{"apple", "pear", "cherry", "plum", "berry"}
	for _, e := range v {
		_ = l.InsertToHead(e)
	}
	p, _ := l.Tail()
	for _, e := range v {
		if r := p.Element(); !reflect.DeepEqual(r, e) {
			t.Errorf("wrong node value: get %+v, expect %+v", r, e)
		}
		p, _ = l.After(p)
	}	
}

func TestBefore(t *testing.T) {
	l := sequences.NewPList[string]()
	v := []string{"apple", "pear", "cherry", "plum", "berry"}
	for _, e := range v {
		_ = l.InsertToTail(e)
	}
	p, _ := l.Head()
	for _, e := range v {
		if r := p.Element(); !reflect.DeepEqual(r, e) {
			t.Errorf("wrong node value: get %+v, expect %+v", r, e)
		}
		p, _ = l.Before(p)
	}	
}

func TestInsertBeforeAfter(t *testing.T) {
	l := sequences.NewPList[string]()
	pos := l.InsertToTail("apple")
	pos, _ = l.InsertBefore("pear", pos)
	_, _ = l.InsertAfter("plum", pos)
	v := []string{"pear", "plum", "apple"}
	pos, _ = l.Tail()
	for _, e := range v {
		if r := pos.Element(); !reflect.DeepEqual(r, e) {
			t.Errorf("wrong InsertBefore/InsertAfter result: receive %+v, expect %+v", r, e)
		}
		pos, _ = l.After(pos)
	}
}

func TestRemoveReplace(t *testing.T) {
	l := sequences.NewPList[string]()
	v := []string{"apple", "pear", "cherry", "berry", "plum"}
	for _, e := range v {
		l.InsertToHead(e)
	}
	pos, _ := l.Tail()
	pos, _ = l.After(pos)
	l.Remove(pos)
	pos, _ = l.Head()
	pos, _ = l.Before(pos)
	_, _ = l.Replace("onion", pos)
	v = []string{"apple", "cherry", "onion", "plum"}
	pos, _ = l.Tail()
	for _, e := range v {
		if r := pos.Element(); !reflect.DeepEqual(r, e) {
			t.Errorf("wrong Remove/Replace result: receive %+v, expect %+v", r, e)
		}
		pos, _ = l.After(pos)
	}
}