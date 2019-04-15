package jsonwebtoken

import (
	"crypto/ecdsa"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
)

// EncodeECDSA encodes and signs a json web token using the claims
// that the json web token should include and an an ECDSA private key.
func EncodeECDSA(claims jwt.Claims, key *ecdsa.PrivateKey) (string, error) {
	token := jwt.New(jwt.SigningMethodES512)
	token.Claims = claims
	signedTokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return signedTokenString, nil
}

// EncodeECDSAFromPemFile encodes and signs a json web token using
// the claims that the json web token should include and an ECDSA
// private key pem container formatted file.
func EncodeECDSAFromPemFile(claims jwt.Claims, file string) (string, error) {
	privateKeyBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	privateKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return "", err
	}

	return EncodeECDSA(claims, privateKey)
}

// DecodeECDSA decodes and verifies a json web token using an ECDSA public key
// and returns a map of the claims held inside the token
func DecodeECDSA(tokenString string, key *ecdsa.PublicKey) (jwt.MapClaims, error) {
	return decodeToken(tokenString, key)
}

// DecodeECDSAFromPemFile decodes and verifies a json web token using an ECDSA
// public key pem container formatted file path and returns a map of the claims
// held inside the token
func DecodeECDSAFromPemFile(tokenString string, file string) (jwt.MapClaims, error) {
	publicKeyBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseECPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, err
	}
	return DecodeECDSA(tokenString, publicKey)
}
