package address

import "github.com/rodrigomonteirof/mastering-bitcoin/internal/secp256k1"

func NewAddress(k secp256k1.PrivateKey) string {
	return "My new address"
}
