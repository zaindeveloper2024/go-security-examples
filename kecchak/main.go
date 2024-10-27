package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/sha3"
)

func main() {
	hash := sha3.New256()
	_, err := hash.Write([]byte("Hello, World!"))
	if err != nil {
		log.Fatalf("failed to write to hash: %v", err)
	}
	sha3 := hash.Sum(nil)
	fmt.Printf("SHA3-256: %x\n", sha3)
}

// fileHash calculates the SHA3-256 hash of a file.
