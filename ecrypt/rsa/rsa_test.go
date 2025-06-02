package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"

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
			t.Errorf("failed to generate private key: %v", err)
			continue
		}
		publicKey := privateKey.PublicKey

		encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, &publicKey, []byte(tc.label+tc.msg))
		if err != nil {
			t.Errorf("failed to ecrypt message %q with label %q: %v", tc.msg, tc.label, err)
			continue
		}

		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encrypted)
		if err != nil {
			t.Errorf("failed to decrypt message %q with label %q: %v", tc.msg, tc.label, err)
			continue
		}

		strDecrypted := string(decrypted)
		label, msg := strDecrypted[:len(tc.label)], strDecrypted[len(tc.label):]
		if label != tc.label || msg != tc.msg {
			t.Errorf("decrypting message %q with label %q: get %q and %q instead", tc.msg, tc.label, msg, label)
		}

	}

}

func TestSignVerifyPSS(t *testing.T) {

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
			t.Errorf("failed to generate private key: %v", err)
			continue
		}
		publicKey := privateKey.PublicKey

		msgHash := sha256.New()
		_, err = msgHash.Write([]byte(tc.msg))
		if err != nil {
			t.Errorf("failed to hash message %q: %v", tc.msg, err)
			continue
		}
		msgHashDigest := msgHash.Sum(nil)

		signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashDigest, nil)
		if err != nil {
			t.Errorf("failed to sign message %q: %v", tc.msg, err)
			continue
		}

		err = rsa.VerifyPSS(&publicKey, crypto.SHA256, msgHashDigest, signature, nil)
		if err != nil {
			t.Errorf("failed to verify message %q: %v", tc.msg, err)
		}

	}

}

func TestMarshalUnmarshalPrivateKey(t *testing.T) {

	tests := []struct {
		bits    int
		pemType string
	}{
		{128, "RSA 128 bit"},
		{256, "RSA 256 bit"},
		{512, "RSA 512 bit"},
		{1024, "RSA 1024 bit"},
		{2048, "RSA 2048 bit"},
		{4096, "RSA 4096 bit"},
	}

	for _, tc := range tests {

		privateKey, err := rsa.GenerateKey(rand.Reader, tc.bits)
		if err != nil {
			t.Errorf("failed to generate %d-bits long private key: %v", tc.bits, err)
		}

		publicKey := privateKey.PublicKey

		// Convert private key to PKCS#1 DER-encoded format
		privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

		// Encode private key to PEM format
		privateKeyPEM := pem.EncodeToMemory(
			&pem.Block{
				Type:  tc.pemType,
				Bytes: privateKeyBytes,
			})

		block, rest := pem.Decode(privateKeyPEM)
		if block == nil {
			t.Errorf("failed to decode PEM byte representration: PEM block is empty")
			continue
		}
		if block.Type != tc.pemType {
			t.Errorf("PEM block type: get %q, expect %q", block.Type, tc.pemType)
			continue
		}
		if len(rest) > 0 {
			t.Errorf("do not expect any rest, get %q", string(rest))
		}

		privateKeyParsed, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			t.Errorf("failed to parse PEM block: %v", err)
		}

		publicKeyParsed := privateKeyParsed.PublicKey
		// if !publicKeyParsed.Equal(publicKey) {
		// 	t.Errorf("%d-bits public key: failed to get same value after marshal/unmarshal", tc.bits)
		// }
		flagModulusEqual := publicKeyParsed.N.Cmp(publicKey.N) == 0
		if publicKeyParsed.E != publicKey.E || !flagModulusEqual {
			t.Errorf("failed to get same public key after marshal/unmarshal\nexponent: get %d, expect %d\nmodulus: %t",
				publicKeyParsed.E,
				publicKey.E,
				flagModulusEqual)
		}

	}

}
