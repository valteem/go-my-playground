package reuse_test

import (
	"testing"

	"github.com/valteem/reuse"
)

func TestEviction(t *testing.T) {

	cache := reuse.NewSimpleCache()

	for i := 0; i < 5; i++ {
		cache.Add(i)
	}

	for i := 0; i < 5; i++ {
		cache.Remove(i)
	}

	for i := 0; i < 5; i++ {
		k, ok := cache.GetOldestEvictedKey()
		if k != i || !ok {
			t.Errorf("GetOldestEvictedKey(): get (%d, %t), expect (%d, true)", k, ok, i)
		}
	}

}

func TestSlicing(t *testing.T) {
	l := []int{0}
	l = l[1:] // magic, index 1 is in fact beyond slice boundaries
	if len(l) > 0 {
		t.Errorf("invalid slice length: %d", len(l))
	}
}
