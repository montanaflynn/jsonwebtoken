package main

import (
	"fmt"
	"log"
	"os"

	"github.com/montanaflynn/jsonwebtoken"
)

var hmacKey = []byte("hunter15")

func main() {
	claims, err := jsonwebtoken.DecodeHMAC(os.Args[1], hmacKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", claims)
}
