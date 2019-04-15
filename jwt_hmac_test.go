package jsonwebtoken

import (
	"crypto/rand"
	"testing"
)

// hold the hmacKey
var hmacKey = make([]byte, 256)

// hold the encoded and signed token
var hmacSignedToken string

func TestEncodeHMAC(t *testing.T) {

	// generate a random hmacKey
	rand.Read(hmacKey)

	var err error
	hmacSignedToken, err = EncodeHMAC(testClaims, hmacKey)
	if err != nil {
		t.Error(err)
	}
}

func TestDecodeHMAC(t *testing.T) {

	claims, err := DecodeHMAC(hmacSignedToken, hmacKey)
	if err != nil {
		t.Error(err)
	}

	if claims["uid"] != 1.0 {
		t.Fail()
	}

	if claims["uem"] != "someone@example.com" {
		t.Fail()
	}

	if claims["iat"] != 1000.0 {
		t.Fail()
	}

	if claims["iss"] != "test" {
		t.Fail()
	}

	if claims["aud"] != "tests" {
		t.Fail()
	}
}
