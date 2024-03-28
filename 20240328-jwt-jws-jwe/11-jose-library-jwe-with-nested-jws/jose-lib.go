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

	signaturePrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	encryptionPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	signedJwt := issueJws(signaturePrivateKey)
	encryptedJwe := encryptJwe(&encryptionPrivateKey.PublicKey, signedJwt)

	fmt.Println(encryptedJwe)

	// receiver part:
	fmt.Println("\nReceiver part:")

	decryptedJwe := decryptJwe(encryptionPrivateKey, encryptedJwe)
	verifyJws(decryptedJwe, &signaturePrivateKey.PublicKey)
}

func issueJws(privateKey *rsa.PrivateKey) string {
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
	return signedJwt
}

func encryptJwe(publicKey *rsa.PublicKey, signedJwt string) string {
	encryptionOpts := (&jose.EncrypterOptions{}).WithContentType("JWT")

	encrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.RSA_OAEP, Key: publicKey}, encryptionOpts)
	if err != nil {
		panic(err)
	}

	encryptedJwe, err := encrypter.Encrypt([]byte(signedJwt))
	if err != nil {
		panic(err)
	}

	serialized, err := encryptedJwe.CompactSerialize()
	if err != nil {
		panic(err)
	}
	return serialized
}

func decryptJwe(privateKey *rsa.PrivateKey, encryptedJwe string) string {
	jwe, err := jose.ParseEncrypted(encryptedJwe, []jose.KeyAlgorithm{jose.RSA_OAEP}, []jose.ContentEncryption{jose.A256GCM})
	if err != nil {
		panic(err)
	}

	decryptedJwe, err := jwe.Decrypt(privateKey)
	if err != nil {
		panic(err)
	}

	return string(decryptedJwe)
}

func verifyJws(signedJwt string, publicKey *rsa.PublicKey) {
	receiverJws, err := jose.ParseSigned(signedJwt, []jose.SignatureAlgorithm{jose.RS256})
	if err != nil {
		panic(err)
	}

	verifiedPayload, err := receiverJws.Verify(publicKey)
	if err != nil {
		fmt.Println("The signature differs - the original message has been modified", err)
	} else {
		fmt.Println("The signature matches - the original message hasn't been modified")
		fmt.Println(string(verifiedPayload))
	}
}
