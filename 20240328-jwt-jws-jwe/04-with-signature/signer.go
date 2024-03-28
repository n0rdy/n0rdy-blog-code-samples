package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	// sender part:
	secretPhase := "goToHellEric"

	part1 := `{
  "typ":"JWT",
  "alg":"HS256"
}`

	part2 := `{
  "sub":"Eric",
  "iat":1710617191,
  "message":"Nice try, Eric, but go play with the kids instead while men are talking here!"
}`

	part1Encoded := encode(part1)
	part2Encoded := encode(part2)
	signature := sign(part1Encoded, part2Encoded, secretPhase)

	encodedMessageWithSignature := part1Encoded + "." + part2Encoded + "." + signature

	fmt.Println(encodedMessageWithSignature)

	// receiver part:
	parts := strings.Split(encodedMessageWithSignature, ".")

	receiverSignature := sign(parts[0], parts[1], secretPhase)
	senderSignature := parts[2]
	if receiverSignature == senderSignature {
		fmt.Println("The signature matches - the original message hasn't been modified")

		part1Decoded := decode(parts[0])
		part2Decoded := decode(parts[1])

		fmt.Println(part1Decoded)
		fmt.Println()
		fmt.Println(part2Decoded)
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

func sign(part1 string, part2 string, secretPhrase string) string {
	textToSign := part1 + "." + part2
	signer := hmac.New(sha256.New, []byte(secretPhrase))

	signer.Write([]byte(textToSign))
	binarySignature := signer.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(binarySignature)
}
