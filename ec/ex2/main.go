package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"log"
)

func main() {
	privkey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	pubkey := privkey.PublicKey

	fmt.Println("Private Key :")
	fmt.Printf("Type: %T Value: %v\n", privkey, privkey)

	fmt.Println("Public Key :")
	fmt.Printf("Type: %T Value: %v\n", pubkey, pubkey)

	// Create a new hash
	h := md5.New()
	io.WriteString(h, "This is a message to be signed and verified by ECDSA!")
	signhash := h.Sum(nil) // convert hash.Hash to byte slice
	fmt.Println("Sign hash: ", signhash)

	r, s, err := ecdsa.Sign(rand.Reader, privkey, signhash)
	if err != nil {
		log.Fatal(err)
	}

	signature := r.Bytes()
	x := s.Bytes() // the x value of the curve (y^2 = x^3 + ax + b)
	signature = append(signature, x...)

	fmt.Printf("Signature : %x\n", signature)

	// Verify
	status := ecdsa.Verify(&pubkey, signhash, r, s)
	fmt.Println(status) // should be true
}
