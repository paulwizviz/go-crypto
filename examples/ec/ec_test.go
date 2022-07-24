package ec

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"reflect"
)

func Example_std() {
	privKey, err := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	fmt.Printf("Type: %T Error: %v\n", privKey, err)

	pubKey := privKey.PublicKey
	fmt.Printf("Type: %T\n", pubKey)

	pubKeyPtr := &privKey.PublicKey
	fmt.Printf("Type: %T\n", pubKeyPtr)

	if reflect.DeepEqual(pubKey, *pubKeyPtr) {
		fmt.Println("Same public key")
	}

	// Output:
	// Type: *ecdsa.PrivateKey Error: <nil>
	// Type: ecdsa.PublicKey
	// Type: *ecdsa.PublicKey
	// Same public key
}
