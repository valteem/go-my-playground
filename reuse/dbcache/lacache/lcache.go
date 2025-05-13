package lcache

import (
	"github.com/valteem/reuse/dbcache"
)

type Cache[K comparable, V any] struct {
	data map[K]V
}

func NewCache[K comparable, V any]() Cache[K, V] {
	return Cache[K, V]{}
}

func (c *Cache[K, V]) Find(key K) (V, bool) {
	v, ok := c.data[key]
	return v, ok
}

func (c *Cache[K, V]) UpdateCacheValue(key K, value V) {
	c.data[key] = value
}

// https://www.reddit.com/r/golang/comments/u3ruav/asserting_generic_type_implements_generic/
var _ dbcache.Cache[int, string] = (*Cache[int, string])(nil)
