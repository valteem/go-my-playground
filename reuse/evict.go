package reuse

type SimpleKeyStorage struct {
	keys       map[int]struct{}
	onEviction func(key int)
}

func (s *SimpleKeyStorage) Remove(key int, onEviction func(key int)) bool {
	if _, ok := s.keys[key]; ok {
		delete(s.keys, key)
		onEviction(key)
		return true
	}
	return false
}

type SimpleCache struct {
	s           *SimpleKeyStorage
	evictedKeys []int
}

func (c *SimpleCache) onEviction(k int) {
	c.evictedKeys = append(c.evictedKeys, k)
}

func NewSimpleCache() *SimpleCache {
	c := &SimpleCache{s: &SimpleKeyStorage{keys: make(map[int]struct{})}}
	c.s.onEviction = c.onEviction // magic
	return c
}

func (c *SimpleCache) Add(key int) {
	c.s.keys[key] = struct{}{}
}

func (c *SimpleCache) Remove(key int) bool {
	removed := c.s.Remove(key, c.s.onEviction)
	return removed
}

func (c *SimpleCache) GetOldestEvictedKey() (int, bool) {
	if len(c.evictedKeys) == 0 {
		return 0, false
	}
	key := c.evictedKeys[0]
	c.evictedKeys = c.evictedKeys[1:]
	return key, true
}
