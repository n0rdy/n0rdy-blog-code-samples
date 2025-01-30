package db

import (
	"20250128-postgres-seq-scan-despite-indexing/common"
	"context"
	"errors"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Repo struct {
	dbPool *pgxpool.Pool
}

func NewRepo(dbPool *pgxpool.Pool) *Repo {
	return &Repo{dbPool: dbPool}
}

func (r *Repo) SelectUsers() ([]common.User, error) {
	rows, err := r.dbPool.Query(
		context.Background(),
		"SELECT id, user_info FROM schema_202501.users",
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var users []common.User
	err = pgxscan.ScanAll(&users, rows)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, nil
}

func (r *Repo) SelectUserBySsn(ssn string) (*common.User, error) {
	now := time.Now()

	var user common.User
	err := r.dbPool.QueryRow(
		context.Background(),
		"SELECT id, user_info FROM schema_202501.users WHERE ssn_hash = crypt($1, ssn_hash)",
		ssn,
	).Scan(&user.Id, &user.UserInfo)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Time taken to get user:", time.Since(now))
	return &user, nil
}

func (r *Repo) InsertUser(user common.NewUserEntity) error {
	_, err := r.dbPool.Exec(
		context.Background(),
		"INSERT INTO schema_202501.users (id, ssn_hash, user_info) VALUES ($1, crypt($2, gen_salt('bf')), $3)",
		user.Id, user.Ssn, user.UserInfo,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *Repo) InsertUsers(users []common.NewUserEntity) error {
	now := time.Now()

	query := "INSERT INTO schema_202501.users (id, ssn_hash, user_info) VALUES "
	var args []interface{}
	for i, user := range users {
		query += fmt.Sprintf("($%d, crypt($%d, gen_salt('bf')), $%d),", i*3+1, i*3+2, i*3+3)
		args = append(args, user.Id, user.Ssn, user.UserInfo)
	}
	query = query[:len(query)-1] // remove the trailing comma

	_, err := r.dbPool.Exec(
		context.Background(),
		query,
		args...,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Time taken to insert users:", time.Since(now))
	return nil
}
