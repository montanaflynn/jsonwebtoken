package jsonwebtoken

import (
	"crypto/rsa"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
)

// EncodeRSA encodes and signs a json web token using the claims
// that the json web token should include and an an RSA private key.
func EncodeRSA(claims jwt.Claims, key *rsa.PrivateKey) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedTokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return signedTokenString, nil
}

// EncodeRSAFromPemFile encodes and signs a json web token using
// the claims that the json web token should include and an RSA
// private key pem container formatted file.
func EncodeRSAFromPemFile(claims jwt.Claims, filePath string) (string, error) {
	privateKeyBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return "", err
	}

	return EncodeRSA(claims, privateKey)
}

// DecodeRSA decodes and verifies a json web token using an RSA public key
// and returns a map of the claims held inside the token.
func DecodeRSA(tokenString string, key *rsa.PublicKey) (jwt.MapClaims, error) {
	return decodeToken(tokenString, key)
}

// DecodeRSAFromPemFile decodes and verifies a json web token using an RSA
// public key pem container formatted file path and returns a map of the
// claims held inside the token.
func DecodeRSAFromPemFile(tokenString string, filePath string) (jwt.MapClaims, error) {
	publicKeyBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, err
	}
	return DecodeRSA(tokenString, publicKey)
}
