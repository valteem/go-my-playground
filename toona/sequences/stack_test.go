package sequences_test

import (
	"testing"

	"my.play.go/toona/sequences"
)

func TestStack(t *testing.T) {
	s := sequences.NewStack[int]()
	(*s).Push(1)
	if v, _ := (*s).Pop(); v != 1 {
		t.Error("wrong Pop() result")
	}
}

func TestAllocStack(t *testing.T) {
	
	s := sequences.NewAllocStack[int](10)
	
	s.Push(1)
	if v, _ := s.Pop(); v != 1 {
		t.Error("wrong Pop() result")
	}

	for i := 1; i <= 10; i++ {
		if e := s.Push(i); e != nil {
			t.Error("wrong Push() result - max stack depth not achieved yet")
		}
	}
	if e := s.Push(11); e == nil {
		t.Error("wrong Push() result - max stack depth achieved error not thrownd")
	}
	
}

func TestMatchDelim(t *testing.T) {

	s1 := `{a+[b+(m+n)]}`
    if sequences.MatchDelim(s1) != true {
		t.Error("wrong MatchDelim result : true expected")
	}
    s2 := "{a+[b+(m+n)]"
	if sequences.MatchDelim(s2) == true {
		t.Error("wrong MatchDelim result : false expected")
	}

}