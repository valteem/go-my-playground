package reuse_test

import (
	"sync"
	"testing"
)

type mergeStr func(string, string) string

func mergeFour(s1, s2 string, f mergeStr) mergeStr {
	fn := func(a1, a2 string) string {
		return s1 + " " + s2 + " " + f(a1, a2)
	}
	return fn
}

func mergeTwo(s1, s2 string, f mergeStr) string {
	return f(s1, s2)
}

func TestAddArgs(t *testing.T) {

	fAdded := mergeFour("simple", "closure", func(a1, a2 string) string {
		return a1 + " " + a2
	})

	if actual, expected := mergeTwo("actually", "works", fAdded), "simple closure actually works"; actual != expected {
		t.Errorf("Get %s, expect %s", actual, expected)
	}

}

func TestSyncWaitGroupClosure(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
		}()
	}

	wg.Wait()

}
