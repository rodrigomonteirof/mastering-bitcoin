package secp256k1

import "math/big"

// curveOrder is N, the order of the generator point G.
// Never pass it as the receiver of a big.Int method: that would mutate it for the whole package.
var curveOrder = mustParseHex("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141")

func mustParseHex(s string) *big.Int {
	v, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic("secp256k1: invalid hex constant " + s)
	}

	return v
}
