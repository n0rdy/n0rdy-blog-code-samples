package main

import (
	"fmt"
	"time"
)

func main() {
	// gets current timestamp:
	current := time.Now().Unix()
	fmt.Println("Current timestamp: ", current)

	// gets a number that is stable within 30 seconds:
	base := current / 30
	fmt.Println("Base: ", base)

	// makes sure it has only 6 digits:
	code := base % 1_000_000

	// adds leading zeros if necessary:
	formattedCode := fmt.Sprintf("%06d", code)
	fmt.Println("Code: ", formattedCode)
}
