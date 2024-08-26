package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

func main() {
	curve := elliptic.P224() //see http://golang.org/pkg/crypto/elliptic/#P256
	privKey, _ := ecdsa.GenerateKey(curve, rand.Reader)
	fmt.Printf("Type: %T Value: %v\n", privKey, privKey)

	pubKey := privKey.PublicKey
	fmt.Printf("Type: %T Value: %v\n", pubKey, pubKey)
}
