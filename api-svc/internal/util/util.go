package util

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateLink(length int) string {
    bytes := make([]byte, (length+3)/4*3)

	// Generate random bytes
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	// Encode the random bytes to base64 URL-safe format
	randomString := base64.URLEncoding.EncodeToString(bytes)

	// Trim any padding characters
	randomString = randomString[:length]

	return randomString
}
