package rsaex

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"reflect"
)

// Generate key using rand Reader from crypto package
func Example_randReader() {

	privkey, err := rsa.GenerateKey(rand.Reader, 1024)
	fmt.Printf("Key Type: %T Error: %v\n", privkey, err)

	pubKey := privkey.PublicKey
	fmt.Printf("Key Type: %T\n", pubKey)

	pubKeyPtr := &privkey.PublicKey
	fmt.Printf("Key Type: %T\n", pubKeyPtr)

	if reflect.DeepEqual(pubKey, *pubKeyPtr) {
		fmt.Println("Same public key")
	}

	// Output:
	// Key Type: *rsa.PrivateKey Error: <nil>
	// Key Type: rsa.PublicKey
	// Key Type: *rsa.PublicKey
	// Same public key

}
