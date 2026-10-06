package main

import (
	"fmt"
	"log"

	"github.com/rodrigomonteirof/mastering-bitcoin/internal/base58"
	"github.com/rodrigomonteirof/mastering-bitcoin/internal/secp256k1"
)

func main() {
	log.Println("generating new private key")

	pk := secp256k1.NewPrivateKey()

	base58 := base58.EncodeCheck(0x80, pk.Value[:])

	fmt.Printf("Private Key: %v\n", base58)
}
