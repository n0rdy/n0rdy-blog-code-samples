package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// 18 + 55 + 1 = 74, so above 72 characters' limit of BCrypt
	userId := randomString(18)
	username := randomString(55)
	password := "super-duper-secure-password"

	combinedString := fmt.Sprintf("%s:%s:%s", userId, username, password)

	combinedHash, err := bcrypt.GenerateFromPassword([]byte(combinedString), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	// let's try to break it
	wrongPassword := "wrong-password"
	wrongCombinedString := fmt.Sprintf("%s:%s:%s", userId, username, wrongPassword)

	err = bcrypt.CompareHashAndPassword(combinedHash, []byte(wrongCombinedString))
	if err != nil {
		fmt.Println("Password is incorrect")
	} else {
		fmt.Println("Password is correct")
	}
}

func randomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length]
}
