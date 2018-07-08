package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func display(key *otp.Key, data []byte) {
	fmt.Printf("Issuer:       %s\n", key.Issuer())
	fmt.Printf("Account Name: %s\n", key.AccountName())
	fmt.Printf("Secret:       %s\n", key.Secret())

	fmt.Println("Writing PNG to qr-code.png...")

	ioutil.WriteFile("qr-code.png", data, 0644)

	fmt.Println()
	fmt.Println("Please add your TOTP to your OTP Application now")
	fmt.Println()
}

func promptForPasscode() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Passcode: ")
	txt, _ := reader.ReadString('\n')
	return txt
}

func main() {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "example.com",
		AccountName: "ryuta@sample.com",
	})
	if err != nil {
		log.Fatalf("failed to generate key: %v", err)
	}

	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		log.Fatalf("failed to create image: %v", err)
	}
	if err := png.Encode(&buf, img); err != nil {
		log.Fatalf("failed to encode to png: %v", err)
	}

	display(key, buf.Bytes())

	fmt.Println("Validating TOTP...")
	passcode := promptForPasscode()
	if totp.Validate(passcode, key.Secret()) {
		fmt.Println("Valid passcode!")
		os.Exit(0)
	} else {
		fmt.Println("Invalid passcode!")
		os.Exit(1)
	}
}
