// Package jsonwebtoken is a set of helper functions for encoding, signing, decoding
// and verifying JSON web tokens built on top of https://github.com/dgrijalva/jwt-go
// with a smaller and simpler API.
//
// Install latest version with:
//
//     go get -u github.com/montanaflynn/jsonwebtoken
//
package jsonwebtoken

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// ErrTokenCouldNotBeParsed is returned when a token could not be parsed
	ErrTokenCouldNotBeParsed = errors.New("token could not be parsed")

	// ErrTokenIsNotJSONWebToken is returned when a token is not a json web token
	ErrTokenIsNotJSONWebToken = errors.New("token is not a json web token")

	// ErrTokenIsNotValid is returned when a token is not valid
	ErrTokenIsNotValid = errors.New("token is not valid")

	// ErrTokenIsNotActiveYet is returned when a token is not active yet
	ErrTokenIsNotActiveYet = errors.New("token is not active yet")

	// ErrTokenIsExpired is returned when a token is expired
	ErrTokenIsExpired = errors.New("token is expired")

	// ErrTokenClaimsCannotBeAsserted is returned when token claims cannot be asserted
	ErrTokenClaimsCannotBeAsserted = errors.New("token claims cannot be asserted")
)

func decodeToken(tokenString string, key interface{}) (jwt.MapClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors == jwt.ValidationErrorMalformed {
				return nil, ErrTokenIsNotJSONWebToken
			} else if ve.Errors == jwt.ValidationErrorExpired {
				return nil, ErrTokenIsExpired
			} else if ve.Errors == jwt.ValidationErrorNotValidYet {
				return nil, ErrTokenIsNotActiveYet
			}
		} else {
			return nil, ErrTokenCouldNotBeParsed
		}
	}

	if !token.Valid {
		return nil, ErrTokenIsNotValid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrTokenClaimsCannotBeAsserted
	}

	return claims, nil
}
