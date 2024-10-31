package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"log"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	message := []byte("Hello, World!")
	fmt.Printf("Original Message: %s\n", message)

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Encrypted Message: %x\n", ciphertext)

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decrypted Message: %s\n", plaintext)
}
