package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func generateSHA256(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func generateFileSHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("error copying file: %s", err)
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func hashWithSalt(input string, salt string) string {
	combined := input + salt
	hash := sha256.New()
	hash.Write([]byte(combined))
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	text := "Hello, World!"
	hashedText := generateSHA256(text)
	fmt.Printf("Text: %s\nHashedText: %s\n", text, hashedText)

	filePath := "go.mod"
	hashedFile, err := generateFileSHA256(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("File: %s\nHashedFile: %s\n", filePath, hashedFile)

	text2 := "Hello, World!"
	salt := "salty"
	hashedText2 := hashWithSalt(text2, salt)
	fmt.Printf("Text: %s\nSalt: %s\nHashedText: %s\n", text2, salt, hashedText2)
}
