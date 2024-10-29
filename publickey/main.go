package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func generateAndSaveKeys() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("failed to generate key pair: %v", err)
	}

	publicKey := &privateKey.PublicKey

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	err = os.WriteFile("private.pem", privateKeyPEM, 0600)
	if err != nil {
		return fmt.Errorf("failed to save private key: %v", err)
	}

	err = os.WriteFile("public.pem", publicKeyPEM, 0644)
	if err != nil {
		return fmt.Errorf("failed to save public key: %v", err)
	}

	return nil
}

func main() {
	err := generateAndSaveKeys()
	if err != nil {
		log.Fatalf("failed to generate and save keys: %v", err)
	}
	fmt.Println("Keys generated and saved successfully")
}
