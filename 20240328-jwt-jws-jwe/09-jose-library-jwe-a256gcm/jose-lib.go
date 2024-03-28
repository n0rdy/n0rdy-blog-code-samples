package main

import (
	"crypto/sha256"
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

	encryptionSecretPhase := "goToHellEric"
	// to make sure the secret is 256 bits long
	hasher := sha256.New()
	hasher.Write([]byte(encryptionSecretPhase))
	encryptionSecret := hasher.Sum(nil)

	payload := Payload{
		Claims: jwt.Claims{
			Subject:  "Craig",
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
		Message: "Hello there, Craig! What's up, buddy?",
	}

	encrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.A256GCMKW, Key: encryptionSecret}, nil)
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

	receiverJwe, err := jose.ParseEncrypted(issuedJwe, []jose.KeyAlgorithm{jose.A256GCMKW}, []jose.ContentEncryption{jose.A256GCM})
	if err != nil {
		panic(err)
	}

	verifiedPayload, err := receiverJwe.Decrypt(encryptionSecret)
	if err != nil {
		fmt.Println("Error decrypting the payload", err)
	} else {
		fmt.Println("The payload has been decrypted successfully")
		fmt.Println(string(verifiedPayload))
	}
}
