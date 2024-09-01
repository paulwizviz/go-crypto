package secp

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
)

func Example_pass() {

	msg := "Hello World"

	privkey, err := btcec.NewPrivateKey()
	if err != nil {
		fmt.Printf("Generate error: %v", err)
	}

	// Signing message (Note for proper Bitcoin project used appropriate hashing)
	h := md5.New()
	io.WriteString(h, msg)
	signhash := h.Sum(nil)

	// Generate signature
	signature := ecdsa.Sign(privkey, signhash)

	status := signature.Verify(signhash, privkey.PubKey())
	fmt.Println(status)

	// Output:
	// true
}

func Example_fail() {

	msg := "Hello World"

	privkey, err := btcec.NewPrivateKey()
	if err != nil {
		fmt.Printf("Generate error: %v", err)
	}

	// Signing message
	h := md5.New()
	io.WriteString(h, msg)
	signhash := h.Sum(nil)

	// Generate signature
	signature := ecdsa.Sign(privkey, signhash)

	privkey1, err := btcec.NewPrivateKey()

	status := signature.Verify(signhash, privkey1.PubKey())
	fmt.Println(status)

	// Output:
	// false
}
