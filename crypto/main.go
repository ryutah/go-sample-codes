package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

type keyPair struct {
	privatekey []byte
	publicKey  []byte
}

func createKeyPairs(bits int) (*keyPair, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}

	return &keyPair{
		privatekey: pem.EncodeToMemory(&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		}),
		publicKey: pem.EncodeToMemory(&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
		}),
	}, nil
}

func main() {
	pair, err := createKeyPairs(2048)
	if err != nil {
		panic(err)
	}
	fmt.Println("Private Key")
	fmt.Printf("%s\n", pair.privatekey)
	fmt.Println("Public Key")
	fmt.Printf("%s\n", pair.publicKey)
}
