package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type keyInfo struct {
	privateKey []byte
	publicKey  []byte
}

func createJWT(privateKey []byte, claim jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	keyBlock, _ := pem.Decode(privateKey)
	if keyBlock == nil {
		return "", errors.New("failed to decode private key")
	}
	key, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		return "", err
	}
	return token.SignedString(key)
}

func createKeys() (*keyInfo, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	bPrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: bPrivateKey,
	}
	privateKeyPem := pem.EncodeToMemory(privateKeyBlock)

	bPublicKey := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: bPublicKey,
	}
	publicKeyPem := pem.EncodeToMemory(publicKeyBlock)

	return &keyInfo{
		privateKey: privateKeyPem,
		publicKey:  publicKeyPem,
	}, nil
}
