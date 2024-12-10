package main

import (
	"fmt"
	"time"
)

func main() {
	var pkForUser1 int64 = 115537846
	var pkForUser2 int64 = 715488689

	codeForUser1 := generateCode(pkForUser1)
	fmt.Println("Code for user 1: ", codeForUser1)

	fmt.Println()

	codeForUser2 := generateCode(pkForUser2)
	fmt.Println("Code for user 2: ", codeForUser2)
}

func generateCode(pk int64) string {
	// gets current timestamp:
	current := time.Now().Unix()
	fmt.Println("Current timestamp: ", current)

	// gets a number that is stable within 30 seconds:
	base := current/30 + pk
	fmt.Println("Base: ", base)

	// makes sure it has only 6 digits:
	code := base % 1_000_000

	// adds leading zeros if necessary:
	return fmt.Sprintf("%06d", code)
}
