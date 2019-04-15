package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

func createKey() (*rsa.PrivateKey, error) {
	reader := rand.Reader
	bitSize := 2048
	return rsa.GenerateKey(reader, bitSize)
}

func savePrivatePEMKey(fileName string, key *rsa.PrivateKey) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	if err != nil {
		return err
	}

	return nil
}

func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) error {
	asn1Bytes, err := x509.MarshalPKIXPublicKey(&pubkey)
	if err != nil {
		return err
	}

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	key, err := createKey()
	if err != nil {
		log.Fatal(err)
	}

	err = savePrivatePEMKey("rsa-private.pem", key)
	if err != nil {
		log.Fatal(err)
	}

	err = savePublicPEMKey("rsa-public.pem", key.PublicKey)
	if err != nil {
		log.Fatal(err)
	}
}
