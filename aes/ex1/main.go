package main

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

func generate64() {
	fmt.Println("-- Generate 64 block cipher --")
	bytesize := 8 // byte
	b := make([]byte, bytesize)
	_, err := rand.Read(b)
	if err != nil {
		log.Printf("Unable to generate random byte slice: %v", err)
		return
	}

	random := hex.EncodeToString(b)
	fmt.Printf("Encoded random string: %s\n", random)

	k, err := hex.DecodeString(random)
	if err != nil {
		log.Printf("Unable to decode random byte string: %v", err)
		return
	}

	fmt.Printf("Byte slice of key: %v\n", k)

	_, err = aes.NewCipher(k)
	if err != nil {
		log.Printf("Unable to generate cipher block: %v", err)
	}

}

func generate128() {
	fmt.Println("-- Generate 128 block cipher --")
	bytesize := 16 // byte
	b := make([]byte, bytesize)
	_, err := rand.Read(b)
	if err != nil {
		log.Printf("Unable to generate random byte slice: %v", err)
		return
	}

	random := hex.EncodeToString(b)
	fmt.Printf("Encoded random string: %s\n", random)

	k, err := hex.DecodeString(random)
	if err != nil {
		log.Printf("Unable to decode random byte string: %v", err)
		return
	}

	fmt.Printf("Byte slice of key: %v\n", k)

	_, err = aes.NewCipher(k)
	if err != nil {
		log.Fatalf("Unable to generate cipher block: %v", err)
	}

}

func generate192() {
	fmt.Println("-- Generate 192 block cipher --")
	bytesize := 24 // byte
	b := make([]byte, bytesize)
	_, err := rand.Read(b)
	if err != nil {
		log.Printf("Unable to generate random byte slice: %v", err)
		return
	}

	random := hex.EncodeToString(b)
	fmt.Printf("Encoded random string: %s\n", random)

	k, err := hex.DecodeString(random)
	if err != nil {
		log.Printf("Unable to decode random byte string: %v", err)
		return
	}

	fmt.Printf("Byte slice of key: %v\n", k)

	_, err = aes.NewCipher(k)
	if err != nil {
		log.Printf("Unable to generate cipher block: %v", err)
	}

}

func generate256() {
	fmt.Println("-- Generate 256 block cipher --")
	bytesize := 32 // byte
	b := make([]byte, bytesize)
	_, err := rand.Read(b)
	if err != nil {
		log.Printf("Unable to generate random byte slice: %v", err)
		return
	}

	random := hex.EncodeToString(b)
	fmt.Printf("Encoded random string: %s\n", random)

	k, err := hex.DecodeString(random)
	if err != nil {
		log.Printf("Unable to decode random byte string: %v", err)
		return
	}

	fmt.Printf("Byte slice of key: %v\n", k)

	_, err = aes.NewCipher(k)
	if err != nil {
		log.Printf("Unable to generate cipher block: %v", err)
	}

}

func generate512() {
	fmt.Println("-- Generate 512 block cipher --")
	bytesize := 64 // byte
	b := make([]byte, bytesize)
	_, err := rand.Read(b)
	if err != nil {
		log.Printf("Unable to generate random byte slice: %v", err)
		return
	}

	random := hex.EncodeToString(b)
	fmt.Printf("Encoded random string: %s\n", random)

	k, err := hex.DecodeString(random)
	if err != nil {
		log.Printf("Unable to decode random byte string: %v", err)
		return
	}

	fmt.Printf("Byte slice of key: %v\n", k)

	_, err = aes.NewCipher(k)
	if err != nil {
		log.Printf("Unable to generate cipher block: %v", err)
	}

}

func main() {
	generate64()
	generate128()
	generate192()
	generate256()
	generate512()
}
