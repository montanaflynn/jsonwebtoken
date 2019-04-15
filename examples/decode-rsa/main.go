package main

import (
	"fmt"
	"log"
	"os"

	"github.com/montanaflynn/jsonwebtoken"
)

func main() {
	claims, err := jsonwebtoken.DecodeRSAFromPemFile(os.Args[1], "./rsa-public.pem")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", claims)
}
