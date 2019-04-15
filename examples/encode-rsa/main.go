package main

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/montanaflynn/jsonwebtoken"
)

func main() {

	// create the claims
	claims := struct {
		UserID    int    `json:"uid"`
		UserEmail string `json:"uem"`
		jwt.StandardClaims
	}{
		UserID:    1,
		UserEmail: "someone@example.com",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			Issuer:    "til-cli",
			Audience:  "internal",
			IssuedAt:  time.Now().Unix(),
			Subject:   "til",
			NotBefore: time.Now().Add(time.Second * 2).Unix(),
		},
	}

	// encode the token
	signedToken, err := jsonwebtoken.EncodeRSAFromPemFile(claims, "./rsa-private.pem")
	if err != nil {
		log.Fatal(err)
	}

	// print the token
	fmt.Println(signedToken)
}
