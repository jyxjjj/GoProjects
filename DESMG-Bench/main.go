package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"github.com/cloudflare/circl/kem/kyber/kyber1024"
	"github.com/cloudflare/circl/kem/mlkem/mlkem1024"
	"golang.org/x/crypto/argon2"
	"io"
	"runtime"
	"time"
)

func randomBytes(length int) []byte {
	b := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic(err)
	}
	return b
}

func main() {
	tsStart := time.Now()
	for i := 0; i < 100; i++ {
		sha256.Sum256(randomBytes(64))
	}
	tsEnd := time.Now()
	tsDuration := tsEnd.Sub(tsStart) / 100
	println("SHA256: ", tsDuration.String())
	tsStart = time.Now()
	for i := 0; i < 100; i++ {
		sha512.Sum512(randomBytes(64))
	}
	tsEnd = time.Now()
	tsDuration = tsEnd.Sub(tsStart) / 100
	println("SHA512: ", tsDuration.String())
	tsStart = time.Now()
	for i := 0; i < 100; i++ {
		block, _ := aes.NewCipher(randomBytes(32))
		gcm, _ := cipher.NewGCM(block)
		nonce := randomBytes(gcm.NonceSize())
		gcm.Seal(nil, nonce, randomBytes(32), randomBytes(32))
	}
	tsEnd = time.Now()
	tsDuration = tsEnd.Sub(tsStart) / 100
	println("AES256GCM: ", tsDuration.String())
	tsStart = time.Now()
	for i := 0; i < 100; i++ {
		scheme := kyber1024.Scheme()
		pk, sk, _ := scheme.GenerateKeyPair()
		ct, _, _ := scheme.Encapsulate(pk)
		_, _ = scheme.Decapsulate(sk, ct)
	}
	tsEnd = time.Now()
	tsDuration = tsEnd.Sub(tsStart) / 100
	println("Kyber1024: ", tsDuration.String())
	tsStart = time.Now()
	for i := 0; i < 100; i++ {
		scheme := mlkem1024.Scheme()
		pk, sk, _ := scheme.GenerateKeyPair()
		ct, _, _ := scheme.Encapsulate(pk)
		_, _ = scheme.Decapsulate(sk, ct)
	}
	tsEnd = time.Now()
	tsDuration = tsEnd.Sub(tsStart) / 100
	println("MLKEM1024: ", tsDuration.String())
	tsStart = time.Now()
	for i := 0; i < 100; i++ {
		argon2.IDKey(
			randomBytes(32),
			randomBytes(32),
			uint32(4),
			uint32(65536),
			uint8(runtime.NumCPU()),
			uint32(64),
		)
	}
	tsEnd = time.Now()
	tsDuration = tsEnd.Sub(tsStart) / 100
	println("Argon2ID: ", tsDuration.String())
}
