package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	encodedMessage := "ewogICJzdWIiOiJDcmFpZyIsCiAgImlhdCI6MTcxMDYxMTk1MywKICAibWVzc2FnZSI6IkhlbGxvIHRoZXJlISIKfQ"
	res, _ := base64.RawURLEncoding.DecodeString(encodedMessage)

	fmt.Println(string(res))
}
