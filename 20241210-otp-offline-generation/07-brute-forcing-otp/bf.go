package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	pk := "JBSWY3DPEB2GQZLSMUQQ"

	code := generateCode(pk)
	fmt.Println("Code: ", code)

	// brute forces the OTP and measures the time it takes to find it in milliseconds:
	start := time.Now()
	for i := 0; i < 1_000_000; i++ {
		asStr := fmt.Sprintf("%06d", i)
		if asStr == code {
			fmt.Println("Found: ", asStr)
			finish := time.Now()
			fmt.Println("Time: ", finish.Sub(start).Milliseconds(), "ms")
			break
		}
	}
}

func generateCode(pk string) string {
	// gets current timestamp:
	current := time.Now().Unix()

	// gets a number that is stable within 30 seconds:
	base := current / 30

	// combines the base with the private key to get a number unique to the user:
	baseWithPk := hash(base, pk)

	// makes sure it is positive:
	absBaseWithPk := int64(math.Abs(float64(baseWithPk)))

	// makes sure it has only 6 digits:
	code := absBaseWithPk % 1_000_000

	// adds leading zeros if necessary:
	return fmt.Sprintf("%06d", code)
}

// converts base and pk into a number
func hash(base int64, pk string) int64 {
	num := base
	for _, char := range pk {
		num = 31*num + int64(char)
	}
	return num
}
