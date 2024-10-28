package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/sha3"
)

func generateFileSHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	hash := sha3.New256()
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("error copying file: %s", err)
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func main() {
	hash := sha3.New256()
	_, err := hash.Write([]byte("Hello, World!"))
	if err != nil {
		log.Fatalf("failed to write to hash: %v", err)
	}
	sha3 := hash.Sum(nil)
	fmt.Printf("SHA3-256: %x\n", sha3)

	filepath := "go.mod"
	hashedFile, err := generateFileSHA256(filepath)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("File: %s\nHashedFile: %s\n", filepath, hashedFile)
}
