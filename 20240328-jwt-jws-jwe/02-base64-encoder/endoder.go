package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	message := `{
  "sub":"Clyde",
  "iat":1710615338,
  "message":"I see that I'm too smart to be friends with a guy like you - don't talk to me, ok?"
}`
	encodedMessage := base64.RawURLEncoding.EncodeToString([]byte(message))

	fmt.Println(encodedMessage)
}
