package main

import (
	"fmt"
	"math/rand"
)

var pkDb = make(map[int]bool)

func main() {
	prForUser := nextPrivateKey()
	fmt.Println("Private key for user: ", prForUser)
}

func nextPrivateKey() int {
	r := randomPrivateKey()
	for pkDb[r] {
		r = randomPrivateKey()
	}
	pkDb[r] = true
	return r
}

// generates random number from 1 000 000 to 999 999 999
func randomPrivateKey() int {
	return rand.Intn(999_999_999-1_000_000) + 1_000_000
}
