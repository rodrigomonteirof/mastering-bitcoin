package wif_test

import (
	"encoding/hex"
	"testing"

	"github.com/rodrigomonteirof/mastering-bitcoin/internal/secp256k1"
	"github.com/rodrigomonteirof/mastering-bitcoin/internal/wif"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name  string
		inHex string
		want  string
	}{
		{"one", "0000000000000000000000000000000000000000000000000000000000000001", "5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreAnchuDf"},
		{"max_valid", "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364140", "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqetqj84qw"},
		{"leading_zero", "00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", "5Hpj8DKG44nNjuTG4mCmPFpnugcNRk98e1K51zfFDWwB4jFa8dT"},
		{"bitcoin_wiki", "0c28fca386c7a227600b2fe50b7cae11ec86d3bf1fbe471be89827e19d72aa1d", "5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ"},
		{"book_example", "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd", "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := hex.DecodeString(tt.inHex)
			if err != nil {
				t.Fatal(err)
			}

			p := secp256k1.PrivateKey{Value: [32]byte(input)}

			got := wif.Encode(p)
			if got != tt.want {
				t.Errorf("EncodeCheck(%s) = %q, want %q", tt.inHex, got, tt.want)
			}
		})
	}
}
