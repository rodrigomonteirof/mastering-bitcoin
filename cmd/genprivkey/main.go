package main

import (
	"fmt"
	"log"

	"github.com/rodrigomonteirof/mastering-bitcoin/internal/secp256k1"
	"github.com/rodrigomonteirof/mastering-bitcoin/internal/wif"
)

func main() {
	log.Println("generating new private key")

	p := secp256k1.NewPrivateKey()

	w := wif.Encode(p)

	fmt.Printf("WIF - Private Key: %v\n", w)

	fmt.Printf("Hex - Private Key: %v\n", p.Hex())

	fmt.Printf("Decimal - Private Key: %v\n", p.Decimal())
}
