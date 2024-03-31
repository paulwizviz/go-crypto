package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
)

// Generate key using rand Reader from crypto package
func main() {
	privkey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Key Type: %T value: %v\n", privkey, privkey)

	pubKey := privkey.PublicKey
	fmt.Printf("Key Type: %T value: %v\n", pubKey, pubKey)
}
