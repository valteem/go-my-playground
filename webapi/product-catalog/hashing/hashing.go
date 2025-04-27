package hashing

type Hasher interface {
	Hash(string) string
}
