package main

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

func main() {
	hash := sha3.New256()
	_, _ = hash.Write([]byte("Hello, World!"))

	sha3 := hash.Sum(nil)

	fmt.Printf("SHA3-256: %x\n", sha3)
}
