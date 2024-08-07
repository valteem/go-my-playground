package hashicorp

import (
	"testing"

	"github.com/hashicorp/golang-lru/v2"
)

const (
	lruLen      = 256
	addCount    = 512
	removeCount = 64
)

func TestLRU(t *testing.T) {

	evictCountExpected := addCount - lruLen
	if evictCountExpected < 0 {
		t.Fatalf("Invalid expected number of evictions: %d", evictCountExpected)
	}
	if removeCount > lruLen {
		t.Fatalf("Invalid remove count: %d", removeCount)
	}

	evictCounter := 0
	onEvicted := func(k int, v int) {
		if k != v {
			t.Fatalf("Evicted key and value should be equal (%d != !%d)", k, v)
		}
		evictCounter++
	}
	l, err := lru.NewWithEvict(lruLen, onEvicted)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	for i := 0; i < addCount; i++ {
		l.Add(i, i)
	}
	if l.Len() != lruLen {
		t.Fatalf("wrong lru length: get %d, expect %d", l.Len(), lruLen)
	}

	if evictCounter != evictCountExpected {
		t.Fatalf("invalid evict count: get %d, expect %d", evictCounter, evictCountExpected)
	}

	for i, k := range l.Keys() {
		if v, ok := l.Get(k); !ok || v != k || v != i+evictCountExpected {
			t.Fatalf("invalid value for key %d: get %d, expect %d", k, v, i+evictCountExpected)
		}
	}

	for i := 0; i < evictCountExpected; i++ {
		if _, ok := l.Get(i); ok {
			t.Fatalf("key %d should have been evicted", i)
		}
	}

	for i := evictCountExpected; i < evictCountExpected+removeCount; i++ {
		l.Remove(i)
		if _, ok := l.Get(i); ok {
			t.Fatalf("key %d should have been deleted", i)
		}
	}

	l.Get(evictCountExpected + removeCount) // this access op moves key to the end of LRU

	for i, k := range l.Keys() {
		if i < len(l.Keys())-1 { // all keys except the last
			if k != i+evictCountExpected+removeCount+1 {
				t.Fatalf("key out of order (i, d): (%d, %d)", i, k)

			}
		} else {
			if k != evictCountExpected+removeCount {
				t.Fatalf("last key out of order (i, d): (%d, %d)", i, k)
			}
		}
	}

	l.Purge()
	if l.Len() != 0 {
		t.Fatalf("invalid LRU length after Purge(): %d", l.Len())
	}
	if v, ok := l.Get(lruLen - removeCount + 1); ok {
		t.Fatalf("Get(%d) - expect nothing, get %d", lruLen-removeCount+1, v)
	}

}

// Test eviction only for fixed-size (size = 1) cache
func TestLRUFixedSizeAdd(t *testing.T) {

	evictCounter := 0
	onEvict := func(k int, v int) {
		evictCounter++
	}

	l, err := lru.NewWithEvict(1, onEvict) // LRU types [K, V] inferred from onEvict() definition
	if err != nil {
		t.Fatalf("error creating new LRU: %v", err)
	}

	if l.Add(1, 1) == true || evictCounter != 0 {
		t.Fatalf("no eviction should have occured")
	}

	if l.Add(2, 2) == false || evictCounter == 0 {
		t.Fatalf("an eviction should have occured")
	}

}

// Contains() vs Get(): does not update `recentness` and eventual eviction
func TestLRUContain(t *testing.T) {

	l, err := lru.New[int, int](2) // nothing to infer [K, V]
	if err != nil {
		t.Fatalf("error creating new LRU: %v", err)
	}

	l.Add(1, 1)
	l.Add(2, 2)

	if !l.Contains(1) {
		t.Fatalf("no eviction if size is not exceeded")
	}

	l.Add(3, 3)
	if l.Contains(1) {
		t.Fatalf("Key `1` should have been evicted")
	}

}

// ContainsOrAdd() does not update recentness
func TestLRUContainsOrAdd(t *testing.T) {

	l, err := lru.New[int, int](2)
	if err != nil {
		t.Fatalf("error creating new LRU: %v", err)
	}

	l.Add(1, 1)
	l.Add(2, 2)

	ok, evicted := l.ContainsOrAdd(1, 11)
	if !ok || evicted {
		t.Fatalf("ContainsOrAdd(1, 11): expect (true, false), get (%t, %t) for (ok, evicted)", ok, evicted)
	}

	l.Add(3, 3)
	ok, evicted = l.ContainsOrAdd(1, 11)
	if ok || !evicted {
		t.Fatalf("ContainsOrAdd(1, 11) after Add(3, 3): expect (false, true), get (%t, %t) for (ok, evicted)", ok, evicted)
	}
}

func TestResize(t *testing.T) {

	evictCounter := 0
	onEviction := func(k, v int) {
		evictCounter++
	}

	l, _ := lru.NewWithEvict(3, onEviction)

	l.Add(1, 1)
	l.Add(2, 2)
	l.Add(3, 3)

	l.Resize(1)

	if evictCounter != 2 {
		t.Fatalf("evict count: get %d, expect 2", evictCounter)
	}

}

func TestRemoveOldest(t *testing.T) {

	evictCounter := 0
	onEviction := func(k, v int) {
		evictCounter++
	}

	l, _ := lru.NewWithEvict(2, onEviction)

	l.Add(1, 1)
	l.Add(2, 2)

	l.RemoveOldest()

	if _, ok := l.Get(1); ok {
		t.Fatalf("key should have been removed")
	}
}

func TestGetOldest(t *testing.T) {

	lastEvictedKey, lastEvictedValue := 0, 0
	onEviction := func(k, v int) {
		lastEvictedKey, lastEvictedValue = k, v
	}

	l, _ := lru.NewWithEvict(2, onEviction)

	l.Add(1, 1)
	l.Add(2, 2)

	key, value, ok := l.GetOldest()
	if key != 1 || value != 1 || !ok {
		t.Fatalf("GetOldest() - ")
	}

	l.Add(3, 3)
	if lastEvictedKey != 1 || lastEvictedValue != 1 {
		t.Fatalf("GetOldest() should have not changed \"recently updated\" status: (%d, %d)", lastEvictedKey, lastEvictedValue)
	}
}
