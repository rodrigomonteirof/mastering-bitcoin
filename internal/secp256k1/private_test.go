package secp256k1_test

import (
	"encoding/hex"
	"testing"

	"github.com/rodrigomonteirof/mastering-bitcoin/internal/secp256k1"
)

func TestNewPrivateKeyIsRandom(t *testing.T) {
	a := secp256k1.NewPrivateKey()
	b := secp256k1.NewPrivateKey()

	if a.Value == b.Value {
		t.Errorf("two generated keys are equal: %x", a.Value)
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		inHex string
		want  bool
	}{
		{"zero", "0000000000000000000000000000000000000000000000000000000000000000", false},
		{"one", "0000000000000000000000000000000000000000000000000000000000000001", true},
		{"n_minus_one", "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364140", true},
		{"n", "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", false},
		{"max_256_bit", "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoded, err := hex.DecodeString(tt.inHex)
			if err != nil {
				t.Fatal(err)
			}
			if len(decoded) != 32 {
				t.Fatalf("inHex has %d bytes, want 32", len(decoded))
			}

			got := secp256k1.IsValid([32]byte(decoded))
			if got != tt.want {
				t.Errorf("IsValid(%s) = got %v, want %v", tt.inHex, got, tt.want)
			}
		})
	}
}
