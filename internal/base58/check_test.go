package base58_test

import (
	"encoding/hex"
	"testing"

	"github.com/rodrigomonteirof/mastering-bitcoin/internal/base58"
)

func TestEncodeCheck(t *testing.T) {
	tests := []struct {
		name    string
		version byte
		inHex   string
		want    string
	}{
		{"genesis_address", 0x00, "62e907b15cbf27d5425399ebf6f0fb50ebb88f18", "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"},
		{"all_zero_payload", 0x00, "0000000000000000000000000000000000000000", "1111111111111111111114oLvT2"},
		{"empty_payload", 0x00, "", "1Wh4bh"},
		{"p2sh", 0x05, "eb15231dfceb60925886b67d065299925915aeb1", "3P823G57hdd4YSuScdQWkpbiiRgmNkzxU5"},
		{"testnet", 0x6f, "eb15231dfceb60925886b67d065299925915aeb1", "n2wxQmfexkjwEPgdD6iJA7T7RtzkiyHwvb"},
		{"wif_uncompressed", 0x80, "0c28fca386c7a227600b2fe50b7cae11ec86d3bf1fbe471be89827e19d72aa1d", "5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ"},
		{"book_example", 0x80, "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd", "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn"},
		{"book_example_compressed", 0x80, "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd01", "KxFC1jmwwCoACiCAWZ3eXa96mBM6tb3TYzGmf6YwgdGWZgawvrtJ"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := hex.DecodeString(tt.inHex)
			if err != nil {
				t.Fatal(err)
			}

			got := base58.EncodeCheck(tt.version, input)
			if got != tt.want {
				t.Errorf("EncodeCheck(%s) = %q, want %q", tt.inHex, got, tt.want)
			}
		})
	}
}
