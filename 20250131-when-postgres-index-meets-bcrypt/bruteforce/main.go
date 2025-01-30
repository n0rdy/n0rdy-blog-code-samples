package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

const (
	ssnLength = 10
)

func main() {
	start := time.Now()

	// initializes the buffer with ASCII zeros
	current := make([]byte, ssnLength)

	// generates all possible combinations
	generateAllPossibleSSNs(current, 0)

	duration := time.Since(start)
	fmt.Printf("Completed in %v\n", duration)
}

// generateAllPossibleSSNs generates all possible SSNs and hashes them
func generateAllPossibleSSNs(current []byte, position int) {
	if position == ssnLength {
		sha256.Sum256(current)
		return
	}

	// converts 0-9 directly to ASCII bytes ('0' = 48 in ASCII)
	for i := byte(48); i < byte(58); i++ {
		current[position] = i
		generateAllPossibleSSNs(current, position+1)
	}
}
