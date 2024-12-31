package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

// Expected result:
//
// 2024/12/31 19:02:55 x509: unknown elliptic curve
// exit status 1

func main() {
	privkey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	der, err := x509.MarshalECPrivateKey(privkey)
	if err != nil {
		log.Fatal(err)
	}

	privPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: der,
	})

	fmt.Println(string(privPEM))
}
