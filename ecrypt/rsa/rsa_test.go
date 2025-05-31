package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	"testing"
)

func TestEncryptDecryptOAEP(t *testing.T) {

	tests := []struct {
		msg   string
		label string
	}{
		{"apples, cherries and pears", "fruits"},
		{"garlics, potatoes and onions", "vegetables"},
		{"tables, chairs and beds", "furniture"},
	}

	for _, tc := range tests {

		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			t.Fatalf("failed to generate private key: %v", err)
		}
		publicKey := privateKey.PublicKey

		encrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey, []byte(tc.msg), []byte(tc.label))
		if err != nil {
			t.Fatalf("failed to ecrypt message %q with label %q: %v", tc.msg, tc.label, err)
		}

		decrypted, err := rsa.DecryptOAEP(sha256.New(), nil, privateKey, encrypted, []byte(tc.label))
		if err != nil {
			t.Fatalf("failed to decrypt message %q with label %q: %v", tc.msg, tc.label, err)
		}

		if msg := string(decrypted); msg != tc.msg {
			t.Errorf("decrypting message %q with label %q: get %q instead", tc.msg, tc.label, msg)
		}

	}

}

func TestEncryptDecryptPKCS1v15(t *testing.T) {

	tests := []struct {
		msg   string
		label string
	}{
		{"apples, cherries and pears", "fruits"},
		{"garlics, potatoes and onions", "vegetables"},
		{"tables, chairs and beds", "furniture"},
	}

	for _, tc := range tests {

		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			t.Fatalf("failed to generate private key: %v", err)
		}
		publicKey := privateKey.PublicKey

		encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, &publicKey, []byte(tc.label+tc.msg))
		if err != nil {
			t.Fatalf("failed to ecrypt message %q with label %q: %v", tc.msg, tc.label, err)
		}

		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encrypted)
		if err != nil {
			t.Fatalf("failed to decrypt message %q with label %q: %v", tc.msg, tc.label, err)
		}

		strDecrypted := string(decrypted)
		label, msg := strDecrypted[:len(tc.label)], strDecrypted[len(tc.label):]
		if label != tc.label || msg != tc.msg {
			t.Errorf("decrypting message %q with label %q: get %q and %q instead", tc.msg, tc.label, msg, label)
		}

	}

}
