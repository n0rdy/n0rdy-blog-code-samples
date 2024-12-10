package main

import (
	"fmt"
	"time"
)

func main() {
	pk := "JBSWY3DPEB2GQZLSMUQQ"

	codeForUser1 := generateCode(pk)
	fmt.Println("Code: ", codeForUser1)
}

func generateCode(pk string) string {
	// gets current timestamp:
	current := time.Now().Unix()
	fmt.Println("Current timestamp: ", current)

	// gets a number that is stable within 30 seconds:
	base := current / 30

	// combines the base with the private key to get a number unique to the user:
	baseWithPk := hash(base, pk)
	fmt.Println("Base: ", baseWithPk)

	// makes sure it has only 6 digits:
	code := baseWithPk % 1_000_000

	// adds leading zeros if necessary:
	return fmt.Sprintf("%06d", code)
}

// converts base and pk into a number
func hash(base int64, pk string) int64 {
	num := base
	for _, char := range pk {
		num = 31*num + int64(char)
	}
	fmt.Println("Hash: ", num)
	return num
}
