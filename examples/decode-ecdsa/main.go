package main

import (
	"fmt"
	"log"
	"os"

	"github.com/montanaflynn/jsonwebtoken"
)

func main() {
	claims, err := jsonwebtoken.DecodeECDSAFromPemFile(os.Args[1], "./ecdsa-public.pem")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", claims)
}
