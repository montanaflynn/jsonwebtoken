# JWT Helper Functions [![][travis-svg]][travis-url] [![][goreport-svg]][goreport-url] [![][godoc-svg]][godoc-url] [![][license-svg]][license-url]

Golang package of helper functions for encoding, signing, decoding and verifying JSON web tokens.

Built on top of [github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go) with a smaller and simpler API.

If you have any suggestions, problems or bug reports please [create an issue](https://github.com/montanaflynn/jsonwebtoken/issues) and I'll do my best to accommodate you. In addition simply starring the repo would show your support for the project and be very much appreciated!

### Install

```
go get -u github.com/montanaflynn/jsonwebtoken
```

### Example Usage

Examples programs using the package for creating and verifying JSON web tokens with `HMAC`, `RSA` and `ECDSA` algorithms and programs for generating a secure `HMAC` key or `RSA` and `ECDSA` public and private key `.pem` files:

algorithm | encoding | decoding | generating keys
----------|----------|----------|----------------
`HMAC`      | [encode hmac jwt](examples/encode-hmac/main.go) | [decode hmac jwt](examples/decode-hmac/main.go) | [generate hmac key](examples/generate-hmac-key/main.go)
`RSA`      | [encode rsa jwt](examples/encode-rsa/main.go) | [decode rsa jwt](examples/decode-rsa/main.go) | [generate rsa keys](examples/generate-rsa-keys/main.go)
`ECDSA`      | [encode ecdsa jwt](examples/encode-ecdsa/main.go) | [decode ecdsa jwt](examples/decode-ecdsa/main.go) | [generate ecdsa keys](examples/generate-ecdsa-keys/main.go)

### Documentation

View the entire package documentation [online at godoc.org](http://godoc.org/github.com/montanaflynn/jsonwebtoken) or offline with `godoc .`

```go
func EncodeHMAC(claims jwt.Claims, key []byte) (string, error) {...}
func DecodeHMAC(tokenString string, key []byte) (jwt.MapClaims, error) {...}

func EncodeRSA(claims jwt.Claims, key *rsa.PrivateKey) (string, error) {...}
func DecodeRSA(tokenString string, key *rsa.PublicKey) (jwt.MapClaims, error) {...}
func EncodeRSAFromPemFile(claims jwt.Claims, file string) (string, error) {...}
func DecodeRSAFromPemFile(tokenString string, file string) (jwt.MapClaims, error) {...}

func EncodeECDSA(claims jwt.Claims, key *ecdsa.PrivateKey) (string, error) {...}
func DecodeECDSA(tokenString string, key *ecdsa.PublicKey) (jwt.MapClaims, error) {...}
func EncodeECDSAFromPemFile(claims jwt.Claims, file string) (string, error) {...}
func DecodeECDSAFromPemFile(tokenString string, file string) (jwt.MapClaims, error) {...}

var ErrTokenCouldNotBeParsed       = errors.New("token could not be parsed")
var ErrTokenIsNotJSONWebToken      = errors.New("token is not a json web token")
var ErrTokenIsNotValid             = errors.New("token is not valid")
var ErrTokenIsNotActiveYet         = errors.New("token is not active yet")
var ErrTokenIsExpired              = errors.New("token is expired")
var ErrTokenClaimsCannotBeAsserted = errors.New("token claims cannot be asserted")
```

[travis-url]: https://travis-ci.org/montanaflynn/jsonwebtoken
[travis-svg]: https://img.shields.io/travis/montanaflynn/jsonwebtoken.svg

[goreport-url]: https://goreportcard.com/report/github.com/montanaflynn/jsonwebtoken
[goreport-svg]: https://goreportcard.com/badge/github.com/montanaflynn/jsonwebtoken

[godoc-url]: https://godoc.org/github.com/montanaflynn/jsonwebtoken
[godoc-svg]: https://godoc.org/github.com/montanaflynn/jsonwebtoken?status.svg

[license-url]: https://github.com/montanaflynn/jsonwebtoken/blob/master/LICENSE
[license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
