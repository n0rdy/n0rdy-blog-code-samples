package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbUrl := "postgres://adminU:adminP@localhost:5432/n0rdyblog"
	dbPool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		panic(err)
	}
	defer dbPool.Close()

	numOfUsers := 5000

	for i := 0; i < numOfUsers; i++ {
		login := randomString(10)
		password := randomString(20)

		fmt.Println("login:", login, "password:", password)

		dbPool.Exec(context.Background(), "INSERT INTO schema_202501.user_credentials(login, password_hash) VALUES($1, crypt($2, gen_salt('bf')))", login, password)
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
