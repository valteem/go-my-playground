package reuse_test

import (
	"testing"

	"github.com/valteem/reuse"
)

func TestCompare(t *testing.T) {
	c1 := reuse.NewCanCompare(1, "a")
	c2 := reuse.NewCanCompare(1, "a")
	if c1 != c2 {
		t.Errorf("Two struct objects should be equal: %+v %+v", c1, c2)
	}

	/*
	n1 := reuse.NewDoNotCompare(1, "a")
	n2 := reuse.NewDoNotCompare(1, "a")
	
	b := n1 == n2
    
	"invalid operation: n1 == n2 (struct containing [0]func() cannot be compared)",
	
	*/ 
}