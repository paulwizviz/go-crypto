package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

func main() {
	privKey, _ := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	fmt.Printf("Type: %T Value: %v\n", privKey, privKey)

	pubKey := privKey.PublicKey
	fmt.Printf("Type: %T Value: %v\n", pubKey, pubKey)
}
