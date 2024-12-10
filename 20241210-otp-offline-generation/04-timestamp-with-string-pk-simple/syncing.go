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
	base := current/30 + hash(pk)
	fmt.Println("Base: ", base)

	// makes sure it has only 6 digits:
	code := base % 1_000_000

	// adds leading zeros if necessary:
	return fmt.Sprintf("%06d", code)
}

// converts the string pk to a number
func hash(pk string) int64 {
	var num int64 = 0
	for _, char := range pk {
		num = 31*num + int64(char)
	}
	return num
}
