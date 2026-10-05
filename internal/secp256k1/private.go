package secp256k1

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
)

const maxKey = "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364140"

type PrivateKey struct {
	Value [32]byte
}

func NewPrivateKey() PrivateKey {
	var k PrivateKey

	for {
		rand.Read(k.Value[:])
		if IsValid(k.Value) {
			return k
		}
	}
}

func IsValid(b [32]byte) bool {
	if b == [32]byte{} {
		return false
	}

	if greaterThanMax(b) {
		return false
	}

	return true
}

func greaterThanMax(b [32]byte) bool {
	decoded, err := hex.DecodeString(maxKey)

	if err != nil || len(b) != 32 {
		panic("secp256k1: invalid constant " + maxKey)
	}

	ret := [32]byte(decoded)

	return bytes.Compare(b[:], ret[:]) > 0
}
