package base58

import (
	"math/big"
	"slices"
)

const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// Encode returns the Base58 representation of b, using the Bitcoin alphabet.
func Encode(b []byte) string {
	var result []byte

	mod := new(big.Int)
	radix := big.NewInt(58)
	n := new(big.Int).SetBytes(b)

	for n.Sign() != 0 {
		n.DivMod(n, radix, mod)
		result = append(result, alphabet[mod.Int64()])
	}

	for range leadingZeros(b) {
		result = append(result, alphabet[0])
	}

	slices.Reverse(result)

	return string(result)
}

func leadingZeros(b []byte) int {
	n := 0
	for _, v := range b {
		if v != 0 {
			break
		}
		n++
	}
	return n
}
