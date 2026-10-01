package base58

import (
	"crypto/sha256"
)

func EncodeCheck(version byte, payload []byte) string {
	ret := make([]byte, 0, 1+len(payload)+4)
	ret = append(ret, version)
	ret = append(ret, payload...)
	ret = append(ret, checksum(ret)...)

	return Encode(ret)
}

func checksum(b []byte) []byte {
	sum1 := sha256.Sum256(b)

	sum2 := sha256.Sum256(sum1[:])

	return sum2[:4]
}
