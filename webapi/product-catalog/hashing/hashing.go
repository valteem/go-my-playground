package hashing

import (
	"crypto/sha1"
	"fmt"
)

type Hasher interface {
	Hash(string) string
}

type SHA1Hasher struct {
	salt string
}

func NewSHA1Hasher(salt string) *SHA1Hasher {
	return &SHA1Hasher{salt: salt}
}

func (h *SHA1Hasher) Hash(input string) string {
	
	hash := sha1.New()
	hash.Write([]byte(input))

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt)))
}
