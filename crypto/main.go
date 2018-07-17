package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

type keyPair struct {
	privatekey []byte
	publicKey  []byte
}

func decryptMessage(privateKey []byte, cipherText []byte) ([]byte, error) {
	key, err := decodeToPrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse key block: %v", err)
	}

	return rsa.DecryptPKCS1v15(rand.Reader, key, cipherText)
}

func encrypteMessage(publicKey []byte, msg []byte) ([]byte, error) {
	key, err := decodeToPublicKey(publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse key block: %v", err)
	}

	return rsa.EncryptPKCS1v15(rand.Reader, key, msg)
}

func decodeToPrivateKey(keybytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(keybytes)
	if block == nil {
		return nil, errors.New("failed to decode private key")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func decodeToPublicKey(keybytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(keybytes)
	if block == nil {
		return nil, errors.New("failed to decode public key")
	}
	return x509.ParsePKCS1PublicKey(block.Bytes)
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
