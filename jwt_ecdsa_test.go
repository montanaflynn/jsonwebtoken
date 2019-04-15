package jsonwebtoken

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"testing"
)

var (
	ecdsaPrivateKey *ecdsa.PrivateKey
	ecdsaPublicKey  *ecdsa.PublicKey
)

func createECDSAKeys(t *testing.T) {
	reader := rand.Reader

	var err error
	ecdsaPrivateKey, err = ecdsa.GenerateKey(elliptic.P521(), reader)
	if err != nil {
		t.Error(err)
	}

	ecdsaPublicKey = &ecdsaPrivateKey.PublicKey
}

var ecdsaSignedToken string

func TestEncodeECDSA(t *testing.T) {

	createECDSAKeys(t)

	var err error
	ecdsaSignedToken, err = EncodeECDSA(testClaims, ecdsaPrivateKey)
	if err != nil {
		t.Error(err)
	}
}

func TestDecodeECDSA(t *testing.T) {
	claims, err := DecodeECDSA(ecdsaSignedToken, ecdsaPublicKey)
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
