package auth_local

import (
	"crypto/hmac"
	"crypto/sha512"
)

const (
	// Size is the size, in bytes, of an authenticated digest.
	Size = 32
	// KeySize is the size, in bytes, of an authentication key.
	KeySize = 32
)

// Sum generates an authenticator for m using a secret key and returns the
// 32-byte digest.
func SumLocal(m []byte, key *[KeySize]byte) *[Size]byte {
	mac := hmac.New(sha512.New, key[:])
	mac.Write(m)
	out := new([Size]byte)
	copy(out[:], mac.Sum(nil)[:Size])
	return out
}