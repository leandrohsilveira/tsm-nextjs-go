package domain

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabasePool(ctx context.Context) (*pgxpool.Pool, error) {
	connString, isSet := os.LookupEnv("DATABASE_URL")
	if !isSet {
		connString = "postgres://app:password@localhost:5432/app?sslmode=disable"
	}

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
