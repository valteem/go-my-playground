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