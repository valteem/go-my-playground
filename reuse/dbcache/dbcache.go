package dbcache

type db[K comparable, V any] interface {
	Find(key K) V
	Update(key K, value V)
}

type Cache[K comparable, V any] interface {
	Find(key K) (V, bool)
	UpdateCacheValue(key K, value V)
	// TODO: add cache eviction
}

type Repository[K comparable, V any] struct {
	cache Cache[K, V]
	db    db[K, V]
}

func NewRepository[K comparable, V any](db db[K, V], cache Cache[K, V]) Repository[K, V] {
	return Repository[K, V]{db: db, cache: cache}
}

func (r *Repository[K, V]) Update(key K, value V) {
	r.db.Update(key, value)
	r.cache.UpdateCacheValue(key, value)
}

func (r *Repository[K, V]) UpdateCacheValue(key K, value V) {
	r.cache.UpdateCacheValue(key, value)
}

func (r *Repository[K, V]) UpdateCache(key K) {
	r.UpdateCacheValue(key, r.db.Find(key))
}

func (r *Repository[K, V]) Find(key K) V {
	if val, ok := r.cache.Find(key); ok {
		return val
	}
	value := r.db.Find(key)
	r.UpdateCacheValue(key, value)
	return value
}
