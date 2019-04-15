package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

func createKey() (*ecdsa.PrivateKey, error) {
	reader := rand.Reader
	return ecdsa.GenerateKey(elliptic.P521(), reader)
}

func savePrivatePEMKey(fileName string, key *ecdsa.PrivateKey) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer outFile.Close()

	privateKeyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	err = pem.Encode(outFile, privateKey)
	if err != nil {
		return err
	}

	return nil
}

func savePublicPEMKey(fileName string, pubkey ecdsa.PublicKey) error {
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

	err = savePrivatePEMKey("ecdsa-private.pem", key)
	if err != nil {
		log.Fatal(err)
	}

	err = savePublicPEMKey("ecdsa-public.pem", key.PublicKey)
	if err != nil {
		log.Fatal(err)
	}
}
