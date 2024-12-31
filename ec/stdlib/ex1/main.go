package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
)

func main() {
	privkey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("D: %[1]T 0x%[1]X\n", privkey.D.Bytes())
	fmt.Printf("X: %[1]T %[1]v\n", privkey.X)
	fmt.Printf("Y: %[1]T %[1]v\n", privkey.Y)
	fmt.Printf("Curve: %[1]T %[1]v\n", privkey.Curve)
}
