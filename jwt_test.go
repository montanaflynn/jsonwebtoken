package jsonwebtoken

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// create the claims
var testClaims = struct {
	UserID    int    `json:"uid"`
	UserEmail string `json:"uem"`
	jwt.StandardClaims
}{
	UserID:    1,
	UserEmail: "someone@example.com",
	StandardClaims: jwt.StandardClaims{
		IssuedAt:  1000,
		NotBefore: 1001,
		ExpiresAt: 7266912725,
		Issuer:    "test",
		Audience:  "tests",
		Subject:   "test",
	},
}
