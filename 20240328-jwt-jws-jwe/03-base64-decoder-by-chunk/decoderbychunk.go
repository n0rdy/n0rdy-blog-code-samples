package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	encodedMessage := "ewogICJ0eXAiOiJKV1QiLAogICJhbGciOiJIUzI1NiIKfQ.ewogICJzdWIiOiJDbHlkZSIsCiAgImlhdCI6MTcxMDYxNzA5MSwKICAibWVzc2FnZSI6IkhleSwgYnVkZHkhIEJvYXJkIGdhbWVzIHRvbmlnaHQhIgp9.ny29zDJjI-QCbihNyPx7hjj0wxpM3E6Isagktf9U-1o"
	chunks := strings.Split(encodedMessage, ".")
	for _, chunk := range chunks {
		res, _ := base64.RawURLEncoding.DecodeString(chunk)
		fmt.Println(string(res) + "\n")
	}
}
