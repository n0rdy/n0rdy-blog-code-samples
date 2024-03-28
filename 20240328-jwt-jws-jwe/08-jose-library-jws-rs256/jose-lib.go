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

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
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

	var signerOpts = jose.SignerOptions{}
	signerOpts.WithType("JWT")

	singer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: privateKey}, &signerOpts)
	if err != nil {
		panic(err)
	}

	signedJwt, err := jwt.Signed(singer).Claims(payload).Serialize()
	if err != nil {
		panic(err)
	}

	fmt.Println(signedJwt)

	// receiver part:
	fmt.Println("\nReceiver part:")

	receiverJws, err := jose.ParseSigned(signedJwt, []jose.SignatureAlgorithm{jose.RS256})
	if err != nil {
		panic(err)
	}

	verifiedPayload, err := receiverJws.Verify(&privateKey.PublicKey)
	if err != nil {
		fmt.Println("The signature differs - the original message has been modified", err)
	} else {
		fmt.Println("The signature matches - the original message hasn't been modified")
		fmt.Println(string(verifiedPayload))
	}
}
