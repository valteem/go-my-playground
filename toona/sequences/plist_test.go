package sequences_test

import (
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