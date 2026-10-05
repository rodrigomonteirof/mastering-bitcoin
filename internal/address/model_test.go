package address_test

import (
	"testing"

	"github.com/rodrigomonteirof/mastering-bitcoin/internal/address"
	"github.com/rodrigomonteirof/mastering-bitcoin/internal/secp256k1"
)

func TestNewAddress(t *testing.T) {
	privateKey := secp256k1.NewPrivateKey()
	result := address.NewAddress(privateKey)

	if result != "My new address" {
		t.Fatalf("got %q, want %q", result, "My new address")
	}
}
