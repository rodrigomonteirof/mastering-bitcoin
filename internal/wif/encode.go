package wif

import (
	"github.com/rodrigomonteirof/mastering-bitcoin/internal/base58"
	"github.com/rodrigomonteirof/mastering-bitcoin/internal/secp256k1"
)

const prefix = 0x80

func Encode(p secp256k1.PrivateKey) string {
	return base58.EncodeCheck(prefix, p.Value[:])
}
