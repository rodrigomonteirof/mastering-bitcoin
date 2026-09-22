package address_test

import (
	"testing"

	"github.com/rodrigomonteirof/mastering-bitcoin/internal/address"
	"github.com/rodrigomonteirof/mastering-bitcoin/internal/secp256k1"
)

func newPrivateKey() secp256k1.PrivateKey {
	return secp256k1.PrivateKey{
		Value: "123456789",
	}
}

func TestNewAddress(t *testing.T) {
	privateKey := newPrivateKey()
	result := address.NewAddress(privateKey)

	if result != "My new address" {
		t.Fatalf("got %q, want %q", result, "My new address")
	}
}
