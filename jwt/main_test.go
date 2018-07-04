package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestEnrypte(t *testing.T) {
	cases := []struct {
		name string
		text string
	}{
		{
			name: "test1",
			text: "hello, world",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			keys, err := createKeys()
			if err != nil {
				t.Fatal(err)
			}

			privateKeyBlock, _ := pem.Decode(keys.privateKey)
			if privateKeyBlock == nil {
				t.Fatal("failed to decode private key")
			}
			privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
			if err != nil {
				t.Fatal(err)
			}

			publicKeyBlock, _ := pem.Decode(keys.publicKey)
			if publicKeyBlock == nil {
				t.Fatal("failed to decode public key")
			}
			publicKey, err := x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
			if err != nil {
				t.Fatal(err)
			}

			cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(c.text))
			if err != nil {
				t.Fatal(err)
			}
			plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
			if err != nil {
				t.Fatal(err)
			}
			if c.text != string(plainText) {
				t.Errorf("want: %s, got: %s", c.text, plainText)
			}
		})
	}
}

func TestSigned(t *testing.T) {
	cases := []struct {
		name string
		text string
	}{
		{
			name: "test1",
			text: "hello, world",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			keys, err := createKeys()
			if err != nil {
				t.Fatal(err)
			}

			privateKeyBlock, _ := pem.Decode(keys.privateKey)
			if privateKeyBlock == nil {
				t.Fatal("failed to decode private key")
			}
			privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
			if err != nil {
				t.Fatal(err)
			}

			publicKeyBlock, _ := pem.Decode(keys.publicKey)
			if publicKeyBlock == nil {
				t.Fatal("failed to decode public key")
			}
			publicKey, err := x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
			if err != nil {
				t.Fatal(err)
			}

			hashed := sha256.Sum256([]byte(c.text))
			digest, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
			if err != nil {
				t.Fatal(err)
			}

			if err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], digest); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestCreateJWT(t *testing.T) {
	cases := []struct {
		name string
		in   jwt.Claims
	}{
		{
			name: "test1",
			in: &jwt.StandardClaims{
				Audience:  "aud",
				ExpiresAt: time.Now().Add(10 * time.Second).Unix(),
				Id:        "id",
				IssuedAt:  time.Now().Unix(),
				Issuer:    "iss",
				NotBefore: time.Now().Unix(),
				Subject:   "sub",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			keys, err := createKeys()
			if err != nil {
				t.Fatal(err)
			}

			tokenstring, err := createJWT(keys.privateKey, c.in)
			t.Log(tokenstring)
			if err != nil {
				t.Fatal(err)
			}

			token, err := jwt.Parse(tokenstring, func(_ *jwt.Token) (interface{}, error) {
				keyBlock, _ := pem.Decode(keys.publicKey)
				if keyBlock == nil {
					return nil, errors.New("failed to decode public key")
				}
				return x509.ParsePKCS1PublicKey(keyBlock.Bytes)
			})

			if err != nil {
				t.Fatal(err)
			}
			if !token.Valid {
				t.Error("invalid token")
			}
		})
	}
}
