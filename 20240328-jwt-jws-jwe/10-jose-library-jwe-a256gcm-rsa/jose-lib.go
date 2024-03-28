package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"

	"time"
)

type Payload struct {
	jwt.Claims
	Message string `json:"message"`
}

func main() {
	// sender part:
	fmt.Println("Sender part:")

	encryptionKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	payload := Payload{
		Claims: jwt.Claims{
			Subject:  "Craig",
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
		Message: "Hello there, Craig! What's up, buddy?",
	}

	encrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.RSA_OAEP, Key: &encryptionKey.PublicKey}, nil)
	if err != nil {
		panic(err)
	}

	issuedJwe, err := jwt.Encrypted(encrypter).Claims(payload).Serialize()
	if err != nil {
		panic(err)
	}

	fmt.Println(issuedJwe)

	// receiver part:
	fmt.Println("\nReceiver part:")

	receiverJwe, err := jose.ParseEncrypted(issuedJwe, []jose.KeyAlgorithm{jose.RSA_OAEP}, []jose.ContentEncryption{jose.A256GCM})
	if err != nil {
		panic(err)
	}

	verifiedPayload, err := receiverJwe.Decrypt(encryptionKey)
	if err != nil {
		fmt.Println("Error decrypting the payload", err)
	} else {
		fmt.Println("The payload has been decrypted successfully")
		fmt.Println(string(verifiedPayload))
	}
}
