package reuse

import (
	"fmt"
	"testing"
)

const (
	numInserts = 100
)

func insert(m map[int]struct{}, n int) {
	for i := 1; i <= n; i++ {
		m[i] = struct{}{}
	}
}

func TestMapInsertOrder(t *testing.T) {
	m1 := map[int]struct{}{}
	insert(m1, numInserts)
	m2 := map[int]struct{}{}
	insert(m2, numInserts)
	fmt.Println(m1)
	fmt.Println(m2)
}