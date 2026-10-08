package secp256k1

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

type PrivateKey struct {
	Value [32]byte
}

func NewPrivateKey() PrivateKey {
	var k PrivateKey

	for {
		rand.Read(k.Value[:])
		if IsValid(k) {
			return k
		}
	}
}

func (p PrivateKey) Decimal() *big.Int {
	d := new(big.Int)

	d.SetBytes(p.Value[:])

	return d
}

func (p PrivateKey) Hex() string {
	return hex.EncodeToString(p.Value[:])
}

func IsValid(p PrivateKey) bool {
	k := p.Decimal()

	return k.Sign() > 0 && k.Cmp(curveOrder) < 0
}
