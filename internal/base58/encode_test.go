package base58_test

import (
	"encoding/hex"
	"testing"

	"github.com/rodrigomonteirof/mastering-bitcoin/internal/base58"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name  string
		inHex string
		want  string
	}{
		{"empty", "", ""},
		{"single_byte", "61", "2g"},
		{"single_byte_low_value", "14", "M"},
		{"three_equal_bytes", "626262", "a3gV"},
		{"three_equal_bytes_next", "636363", "aPEr"},
		{"ascii_long_string", "73696d706c792061206c6f6e6720737472696e67", "2cFupjhnEsSn59qHXstmK2ffpLv2"},
		{"five_bytes", "516b6fcd0f", "ABnLTmg"},
		{"interior_zero_byte", "bf4f89001e670274dd", "3SEo3LWLoPntC"},
		{"four_bytes", "572e4794", "3EFU7m"},
		{"ten_bytes", "ecac89cad93923c02321", "EJDM8drfXA6uyA"},
		{"leading_high_nibble_zero", "10c8511e", "Rt5zm"},
		{"all_zero_bytes", "00000000000000000000", "1111111111"},
		{"mainnet_address", "00eb15231dfceb60925886b67d065299925915aeb172c06647", "1NS17iag9jJgTHD1VXjvLCEnZuQ3rJDE9L"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := hex.DecodeString(tt.inHex)
			if err != nil {
				t.Fatal(err)
			}

			got := base58.Encode(input)
			if got != tt.want {
				t.Errorf("Encode(%s) = %q, want %q", tt.inHex, got, tt.want)
			}
		})
	}
}
