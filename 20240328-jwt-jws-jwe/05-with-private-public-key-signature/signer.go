package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	// sender part:
	fmt.Println("Sender part:")

	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey

	header := `{
  "typ":"JWT",
  "alg":"RS256"
}`

	payload := `{
  "sub":"Clyde",
  "iat":1710617191,
  "message":"Hey, buddy! Board games tonight!"
}`

	headerEncoded := encode(header)
	payloadEncoded := encode(payload)
	signature := sign(headerEncoded, payloadEncoded, privateKey)

	encodedMessageWithSignature := headerEncoded + "." + payloadEncoded + "." + signature

	fmt.Println(encodedMessageWithSignature)

	// receiver part:
	fmt.Println("\nReceiver part:")

	parts := strings.Split(encodedMessageWithSignature, ".")
	senderSignature := parts[2]
	if verify(parts[0], parts[1], senderSignature, publicKey) {
		fmt.Println("The signature is valid - the original message hasn't been modified")

		headerDecoded := decode(parts[0])
		payloadDecoded := decode(parts[1])

		fmt.Println(headerDecoded)
		fmt.Println()
		fmt.Println(payloadDecoded)
	} else {
		fmt.Println("The signature differs - the original message has been modified")
	}
}

func encode(text string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(text))
}

func decode(text string) string {
	decoded, _ := base64.RawURLEncoding.DecodeString(text)
	return string(decoded)
}

func sign(header string, payload string, privateKey *rsa.PrivateKey) string {
	textToSign := header + "." + payload
	hashed := sha256.Sum256([]byte(textToSign))

	binarySignature, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	return base64.RawURLEncoding.EncodeToString(binarySignature)
}

func verify(header string, payload string, signature string, publicKey *rsa.PublicKey) bool {
	textToSign := header + "." + payload
	hashed := sha256.Sum256([]byte(textToSign))

	binarySignature, _ := base64.RawURLEncoding.DecodeString(signature)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], binarySignature)
	return err == nil
}
