package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

func main() {
	// sender part:
	fmt.Println("Sender part:")

	signingSecretPhase := "goToHellEric"
	encryptionSecretPhase := "keepItSecretKeepItSafe"

	jweHeader := `{
  "typ":"JWE",	
  "alg":"dir",
  "enc":"A256GCM",	
}`

	payload := `{
  "sub":"Craig",
  "iat":1710617455,
  "message":"I think we should include elves in our videogame!"
}`

	headersEncoded := encode(jweHeader)
	payloadEncoded := encode(payload)
	payloadEncrypted, _ := encrypt(payloadEncoded, encryptionSecretPhase)
	signature := sign(headersEncoded, payloadEncrypted, signingSecretPhase)

	encodedMessageWithSignature := headersEncoded + ".." + payloadEncrypted + "." + signature

	fmt.Println(encodedMessageWithSignature)

	// receiver part:
	fmt.Println("\nReceiver part:")

	parts := strings.Split(encodedMessageWithSignature, ".")
	headersReceived := parts[0]
	payloadReceived := parts[2] + "." + parts[3]
	signatureReceived := parts[4]

	receiverSignature := sign(headersReceived, payloadReceived, signingSecretPhase)
	if receiverSignature == signatureReceived {
		fmt.Println("The signature matches - the original message hasn't been modified")

		headersDecoded := decode(headersReceived)
		payloadDecrypted, _ := decrypt(payloadReceived, encryptionSecretPhase)
		payloadDecoded := decode(payloadDecrypted)

		fmt.Println(headersDecoded)
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

func encrypt(payload string, secretPhrase string) (string, error) {
	hash := sha256.Sum256([]byte(secretPhrase))
	secretKey := hash[:]

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	encodedNonce := base64.RawURLEncoding.EncodeToString(nonce)

	ciphertext := gcm.Seal(nil, nonce, []byte(payload), nil)
	encodedCiphertext := base64.RawURLEncoding.EncodeToString(ciphertext)
	return fmt.Sprintf("%s.%s", encodedNonce, encodedCiphertext), nil
}

func decrypt(encryptedPayload string, secretPhrase string) (string, error) {
	hash := sha256.Sum256([]byte(secretPhrase))
	secretKey := hash[:]

	// The encryptedPayload includes the IV and the ciphertext, separated by a dot.
	parts := strings.Split(encryptedPayload, ".")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid encrypted payload format")
	}
	encodedNonce, encodedCiphertext := parts[0], parts[1]

	nonce, err := base64.RawURLEncoding.DecodeString(encodedNonce)
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.RawURLEncoding.DecodeString(encodedCiphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil

}

func sign(part1 string, part2 string, secretPhrase string) string {
	textToSign := part1 + "." + part2
	signer := hmac.New(sha256.New, []byte(secretPhrase))

	signer.Write([]byte(textToSign))
	binarySignature := signer.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(binarySignature)
}
