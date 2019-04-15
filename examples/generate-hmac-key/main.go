package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func secureString(byteLength int) (string, error) {
	bytes := make([]byte, byteLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes), err
}

func main() {
	key, err := secureString(64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(key)
}
