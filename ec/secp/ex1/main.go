package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privkey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("D: %[1]T 0x%[1]X\n", privkey.D.Bytes()) // []uint8
	fmt.Printf("X: %[1]T %[1]v\n", privkey.X)           // *big.Int
	fmt.Printf("Y: %[1]T %[1]v\n", privkey.Y)           // *big.Int
	fmt.Printf("Curve: %[1]T %[1]v\n", privkey.Curve)

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	kspath := path.Join(pwd, "tmp")
	if _, err := os.Stat(kspath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(kspath, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	passphrase := "passphrase"

	// Serialised into key store file
	ks := keystore.NewKeyStore(kspath, keystore.StandardScryptN, keystore.StandardScryptP)
	acct, err := ks.ImportECDSA(privkey, passphrase)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%[1]T %[1]v\n", acct)

	// Deserialise from keystore file
	keystoreJSON, err := os.ReadFile(acct.URL.Path)
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(keystoreJSON, passphrase)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%[1]T %[1]v\n", key.PrivateKey)

	fmt.Printf("%v", privkey.D.Cmp(key.PrivateKey.D) == 0) // true
}
