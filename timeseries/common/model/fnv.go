package model

// https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function

const (
	offset64 = 14695981039346656037
	prime64  = 1099511628211
)

// Initializes a new fnv64a hash value
func hashNew() uint64 {
	return offset64
}

// Adds a string to a fnv64a hash value, returns the updated hash value
func hashAdd(h uint64, s string) uint64 {
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= prime64
	}
	return h
}

// Adds a byte to a fnv64a hash value, returns the updated hash value
func hashAddByte(h uint64, b byte) uint64 {
	h ^= uint64(b)
	h *= prime64
	return h
}
