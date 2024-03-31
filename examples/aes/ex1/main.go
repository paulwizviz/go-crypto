package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

func encrypt(key string, plaintext string) (string, error) {
	k, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	pt := []byte(plaintext)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}
	ct := gcm.Seal(nonce, nonce, pt, nil)
	return fmt.Sprintf("%x", ct), nil
}

func decrypt(ciphertext string, key string) (string, error) {
	k, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	enc, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(k)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	noncesize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ct := enc[:noncesize], enc[noncesize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ct, nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plaintext), nil
}

func main() {

	// Generate random key
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	key := hex.EncodeToString(b)
	plaintext := "This is a test"

	ct, err := encrypt(key, plaintext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encrypted text: ", ct)

	plaintext, err = decrypt(ct, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Plain text: ", plaintext)
}
