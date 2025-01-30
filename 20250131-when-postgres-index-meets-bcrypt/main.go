package main

import (
	"20250128-postgres-seq-scan-despite-indexing/api"
	"20250128-postgres-seq-scan-despite-indexing/db"
	"20250128-postgres-seq-scan-despite-indexing/service"
	"20250128-postgres-seq-scan-despite-indexing/utils"
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"strings"
)

func main() {
	dbUrl := "postgres://adminU:adminP@localhost:5432/n0rdyblog"
	runMigrations(dbUrl)

	dbPool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		panic(err)
	}
	defer dbPool.Close()

	repo := db.NewRepo(dbPool)
	srv := service.NewService(repo)
	router := api.NewApiRouter(srv)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: router.NewRouter(),
	}

	fillDb(repo)

	err = server.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("server shutdown")
		} else {
			fmt.Println("server failed")
			panic(err)
		}
	}
}

func runMigrations(dbUrl string) {
	m, err := migrate.New(
		"file://db/migrations",
		"pgx5://"+stripPostgresPrefix(dbUrl),
	)
	if err != nil {
		panic(fmt.Errorf("failed to create migrate instance: %w", err))
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(fmt.Errorf("failed to run migrations: %w", err))
	}
}

func stripPostgresPrefix(dbUrl string) string {
	return strings.TrimPrefix(dbUrl, "postgres://")
}

func fillDb(repo *db.Repo) {
	users, err := repo.SelectUsers()
	if err != nil {
		panic(err)
	}
	if len(users) > 0 {
		fmt.Println("users already exist, skipping db fill")
		return
	}

	fmt.Println("generating random users...")
	newUsers := utils.GenRandomUsers(5000)

	fmt.Println("inserting random users...")
	err = repo.InsertUsers(newUsers)
	if err != nil {
		panic(err)
	}
	fmt.Println("inserted random users")
}
