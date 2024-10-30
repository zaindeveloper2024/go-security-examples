package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func createHMAC(message, key string) string {
	h := hmac.New(sha256.New, []byte(key))

	h.Write([]byte(message))

	return hex.EncodeToString(h.Sum(nil))
}

func validateHMAC(message, key, expectedMAC string) bool {
	mac := createHMAC(message, key)

	return hmac.Equal([]byte(mac), []byte(expectedMAC))
}

func main() {
	message := "Hello, World!"
	key := "secret key"

	messageMAC := createHMAC(message, key)
	fmt.Printf("Message: %s\n", message)

	isValid := validateHMAC(message, key, messageMAC)
	fmt.Printf("Message MAC: %v\n", isValid)

	fakeHMAC := "fake hmac"
	isInvalid := validateHMAC(message, key, fakeHMAC)
	fmt.Printf("Fake HMAC: %v\n", isInvalid)
}
