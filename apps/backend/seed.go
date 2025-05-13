package main

import (
	"context"
	"fmt"
	"os"
	"tsm/database"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Seed(ctx context.Context, pool *pgxpool.Pool) error {
	activate := ShouldSeed()

	if !activate {
		return nil
	}

	username, isUsernameSet := os.LookupEnv("ADMIN_USERNAME")
	if !isUsernameSet {
		username = "admin@email.com"
	}

	password, isPasswordSet := os.LookupEnv("ADMIN_PASSWORD")
	if !isPasswordSet {
		password = "123456"
	}

	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	queries := database.New(conn)

	_, err = queries.GetUserByEmail(ctx, username)

	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	if err == nil {
		fmt.Printf("Admin user %s already exists, aborting...", username)
		return nil
	}

	_, err = queries.CreateUser(ctx, database.CreateUserParams{
		Name:     username,
		Email:    username,
		Password: pgtype.Text{String: password},
		Role:     database.UserRoleAdminUser,
	})

	if err == nil {
		fmt.Printf("Admin user %s created", username)
	}

	return err
}

func ShouldSeed() bool {
	args := os.Args[1:]

	activate := false
	for _, arg := range args {
		if arg == "--seed" {
			activate = true
		}
	}

	return activate
}
