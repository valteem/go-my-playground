package orderedmap

import (
	"sync"
	"my.play.go/toona/collections/slice"
)

type OrderedMap[K comparable, V any] struct{
	mu sync.Mutex
	s []K
	m map[K]V
}

func New[K comparable, V any]() OrderedMap[K, V] {
	return OrderedMap[K, V]{
		s: make([]K, 0),
		m: make(map[K]V),
	}
}

func (o *OrderedMap[K, V]) Insert(key K, value V) {
	o.mu.Lock()
	if !slice.Contains(o.s, key) { // only new keys are added to slice, otherwise previous value for existing key is overwritten
		o.s = append(o.s, key)
	}
	o.m[key] = value
	o.mu.Unlock()
}

func (o *OrderedMap[K, V]) Keys() []K {
	o.mu.Lock()
	s := make([]K, len(o.s))
	copy(s, o.s)
	o.mu.Unlock()
	return s
}

func (o *OrderedMap[K, V]) Values() []V {
	o.mu.Lock()
	s := []V{}
	for _, v := range o.s {
		s = append(s, o.m[v])
	}
	o.mu.Unlock()
	return s
}

