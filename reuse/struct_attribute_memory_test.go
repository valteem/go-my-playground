package reuse_test

import (
	"fmt"
	"testing"
	"github.com/valteem/reuse"
)

func TestStructAttributeMemory(t *testing.T) {

	n := reuse.NewNode(1)
// Go creates a sort of 'deep copy' of {n} while assigning it as attribute value for {tree}
	tree := reuse.NewTree(n)
	tree.Root.Key = 2

	if n.Key != 1 {
		t.Errorf("wrong attribute value")
	}

	fmt.Printf("%p\n", &n)
	fmt.Printf("%p\n", &(tree.Root))
}