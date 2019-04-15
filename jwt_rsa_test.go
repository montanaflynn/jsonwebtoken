package jsonwebtoken

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"testing"
)

var (
	rsaPrivateKey *rsa.PrivateKey
	rsaPublicKey  *rsa.PublicKey
)

func createRSAKeys(t *testing.T) {
	reader := rand.Reader
	bitSize := 2048

	var err error
	rsaPrivateKey, err = rsa.GenerateKey(reader, bitSize)
	if err != nil {
		t.Error(err)
	}

	rsaPublicKey = &rsaPrivateKey.PublicKey
}

var rsaSignedToken string

func TestEncodeRSA(t *testing.T) {

	createRSAKeys(t)

	var err error
	rsaSignedToken, err = EncodeRSA(testClaims, rsaPrivateKey)
	if err != nil {
		t.Error(err)
	}
}

func TestDecodeRSA(t *testing.T) {
	claims, err := DecodeRSA(rsaSignedToken, rsaPublicKey)
	if err != nil {
		log.Fatal(err)
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
