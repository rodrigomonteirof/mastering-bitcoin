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

func TestHex(t *testing.T) {
	tests := []struct {
		name string
		in   [32]byte
		want string
	}{
		{"one", [32]byte{31: 0x01}, "0000000000000000000000000000000000000000000000000000000000000001"},
		{"first_byte", [32]byte{0: 0xff}, "ff00000000000000000000000000000000000000000000000000000000000000"},
		{"first_and_last", [32]byte{0: 0xab, 31: 0xcd}, "ab000000000000000000000000000000000000000000000000000000000000cd"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := secp256k1.PrivateKey{Value: tt.in}

			got := k.Hex()
			if got != tt.want {
				t.Errorf("Hex() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDecimal(t *testing.T) {
	tests := []struct {
		name  string
		inHex string
		want  string
	}{
		{"one", "0000000000000000000000000000000000000000000000000000000000000001", "1"},
		{"book_example", "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd", "13840170145645816737842251482747434280357113762558403558088249138233286766301"},
		{"n_minus_one", "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364140", "115792089237316195423570985008687907852837564279074904382605163141518161494336"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := keyFromHex(t, tt.inHex)

			got := k.Decimal().String()
			if got != tt.want {
				t.Errorf("Decimal() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestDecimalReturnsCopy(t *testing.T) {
	k := keyFromHex(t, "0000000000000000000000000000000000000000000000000000000000000001")

	d := k.Decimal()
	d.SetInt64(42)

	if got := k.Decimal().String(); got != "1" {
		t.Errorf("mutating Decimal() result changed the key: got %s, want 1", got)
	}
}

func keyFromHex(t *testing.T, s string) secp256k1.PrivateKey {
	t.Helper()

	decoded, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	if len(decoded) != 32 {
		t.Fatalf("hex has %d bytes, want 32", len(decoded))
	}

	return secp256k1.PrivateKey{Value: [32]byte(decoded)}
}
