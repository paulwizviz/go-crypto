package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	// 1. Generate an ECDSA private key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Marshal the private key to DER format (PKCS#8)
	derBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Choose a password (use a strong, randomly generated password in real applications)
	password := []byte("mysecretpassword")

	// 4. Generate a salt
	salt := make([]byte, 8)
	if _, err := rand.Read(salt); err != nil {
		log.Fatal(err)
	}

	// 5. Derive a key using PBKDF2
	key := pbkdf2.Key(password, salt, 10000, 32, sha256.New)

	// 6. Create an AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// 7. Create a GCM cipher
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	// 8. Create a nonce
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		log.Fatal(err)
	}

	// 9. Encrypt the private key
	ciphertext := aesgcm.Seal(nil, nonce, derBytes, nil)

	fmt.Println("Salt Length: ", len(salt))
	fmt.Println("Nonce: ", len(nonce))
	fmt.Println("Cipher text: ", len(ciphertext))

	// 10. Create a PEM block
	pemBlock := &pem.Block{
		Type:  "ENCRYPTED PRIVATE KEY",
		Bytes: append(salt, append(nonce, ciphertext...)...), // Store salt, nonce, and ciphertext
	}

	// 11. Encode the PEM block
	pemBytes := pem.EncodeToMemory(pemBlock)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	tmpPath := path.Join(wd, "tmp")
	if _, err := os.Stat(tmpPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(tmpPath, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 12. Write the PEM bytes to a file
	err = os.WriteFile(path.Join(tmpPath, "encrypted_private.pem"), pemBytes, 0600)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encrypted private key saved to encrypted_private.pem")

	// --- Decryption Example ---
	pemData, err := os.ReadFile(path.Join(tmpPath, "encrypted_private.pem"))
	if err != nil {
		log.Fatal(err)
	}

	blockPem, _ := pem.Decode(pemData)
	if blockPem == nil || blockPem.Type != "ENCRYPTED PRIVATE KEY" {
		log.Fatal(fmt.Errorf("failed to decode PEM block containing encrypted private key"))
	}

	data := blockPem.Bytes
	saltDec := data[:8]
	data = data[8:]
	nonceDec := data[:12]
	ciphertextDec := data[12:]

	keyDec := pbkdf2.Key(password, saltDec, 10000, 32, sha256.New)

	blockDec, err := aes.NewCipher(keyDec)
	if err != nil {
		log.Fatal(err)
	}

	aesgcmDec, err := cipher.NewGCM(blockDec)
	if err != nil {
		log.Fatal(err)
	}

	plaintextDec, err := aesgcmDec.Open(nil, nonceDec, ciphertextDec, nil)
	if err != nil {
		log.Fatal(err)
	}

	keyInterface, err := x509.ParsePKCS8PrivateKey(plaintextDec)
	if err != nil {
		log.Fatal(err)
	}

	privateKeyLoaded, ok := keyInterface.(*ecdsa.PrivateKey)
	if !ok {
		log.Fatal(fmt.Errorf("failed to parse decrypted private key"))
	}

	fmt.Println("Private Key Loaded successfully")
	fmt.Printf("%[1]T %[1]v", privateKeyLoaded)
}
