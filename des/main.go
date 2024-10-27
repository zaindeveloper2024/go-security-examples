package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func PKCS5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func PKCS5UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func generateIV() ([]byte, error) {
	iv := make([]byte, des.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, fmt.Errorf("failed to generate IV: %v", err)
	}
	return iv, nil
}

func EncryptDES(data []byte, key []byte) (string, string, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return "", "", fmt.Errorf("failed to create cipher: %v", err)
	}

	iv, err := generateIV()
	if err != nil {
		return "", "", fmt.Errorf("failed to generate IV: %v", err)
	}

	paddedData := PKCS5Padding(data, block.BlockSize())

	mode := cipher.NewCBCEncrypter(block, iv)

	encrypted := make([]byte, len(paddedData))
	mode.CryptBlocks(encrypted, paddedData)

	encryptedBase64 := base64.StdEncoding.EncodeToString(encrypted)
	ivBase64 := base64.StdEncoding.EncodeToString(iv)

	return encryptedBase64, ivBase64, nil
}

func EncryptDESWithCTR(data []byte, key []byte) (string, string, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return "", "", fmt.Errorf("failed to create cipher: %v", err)
	}

	iv, err := generateIV()
	if err != nil {
		return "", "", fmt.Errorf("failed to generate IV: %v", err)
	}

	paddedData := PKCS5Padding(data, block.BlockSize())

	mode := cipher.NewCTR(block, iv)

	encrypted := make([]byte, len(paddedData))
	mode.XORKeyStream(encrypted, paddedData)

	encryptedBase64 := base64.StdEncoding.EncodeToString(encrypted)
	ivBase64 := base64.StdEncoding.EncodeToString(iv)

	return encryptedBase64, ivBase64, nil
}

func DecryptDES(encryptedBase64 string, ivBase64 string, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %v", err)
	}

	encrypted, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %v", err)
	}

	iv, err := base64.StdEncoding.DecodeString(ivBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode IV: %v", err)
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	decrypted := make([]byte, len(encrypted))
	mode.CryptBlocks(decrypted, encrypted)

	decrypted = PKCS5UnPadding(decrypted)

	return decrypted, nil
}

func DecryptDESWithCTR(encryptedBase64 string, ivBase64 string, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %v", err)
	}

	encrypted, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %v", err)
	}

	iv, err := base64.StdEncoding.DecodeString(ivBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode IV: %v", err)
	}

	mode := cipher.NewCTR(block, iv)

	decrypted := make([]byte, len(encrypted))
	mode.XORKeyStream(decrypted, encrypted)

	decrypted = PKCS5UnPadding(decrypted)

	return decrypted, nil
}

func main() {
	key := []byte("12345678")

	originalText := "Hello, world!"
	fmt.Printf("Original: %s\n", originalText)

	encrypted, iv, err := EncryptDES([]byte(originalText), key)
	if err != nil {
		log.Fatalf("Encryption failed: %v", err)
	}
	fmt.Printf("Encrypted (base64): %s\n", encrypted)

	decrypted, err := DecryptDES(encrypted, iv, key)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}
	fmt.Printf("Decrypted: %s\n", string(decrypted))

	fmt.Println("---------> CTR MODE <---------")

	encrypted, iv, err = EncryptDESWithCTR([]byte(originalText), key)
	if err != nil {
		log.Fatalf("Encryption failed: %v", err)
	}
	fmt.Println("Encrypted (base64):", encrypted)

	decrypted, err = DecryptDESWithCTR(encrypted, iv, key)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}
	fmt.Printf("Decrypted: %s\n", string(decrypted))
}
