package jsonwebtoken

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// EncodeHMAC encodes and signs a json web token using the claims
// that the json web token should include and an HMAC secret key.
func EncodeHMAC(claims jwt.Claims, key []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return signedTokenString, nil
}

// DecodeHMAC decodes and verifies a json web token using an HMAC
// secret key and returns a map of the claims held inside the token.
func DecodeHMAC(tokenString string, key []byte) (jwt.MapClaims, error) {
	return decodeToken(tokenString, key)
}
