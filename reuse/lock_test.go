package reuse_test

import (
	"sync"
	"testing"
)

func countUniqueValues(input []int) map[int]int {
	m := make(map[int]int, 0)
	for _, v := range input {
		m[v]++
	}
	return m
}

func addToStore(store *[]int, l *sync.Mutex, w *sync.WaitGroup, value int) {
	l.Lock()
	defer l.Unlock()
	*store = append(*store, value)
	w.Done()
}

func TestLock(t *testing.T) {
	l := sync.Mutex{}
	size := 100 * 100 * 10
	store := make([]int, size*2)
	w := sync.WaitGroup{}
	w.Add(2 * size)
	for i := 0; i < size; i++ {
		go addToStore(&store, &l, &w, 1)
		go addToStore(&store, &l, &w, 2)
	}
	w.Wait()
	count := countUniqueValues(store)
	if count[1] != size || count[2] != size {
		t.Errorf("expect equal count of 1 and 2, get %d, %d", count[1], count[2])
	}
}
