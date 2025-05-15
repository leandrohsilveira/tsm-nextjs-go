package domain

import (
	"context"
	"os"
	"tsm/database"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type DatabasePool struct {
	pool   *pgxpool.Pool
	logger echo.Logger
}

var ErrNoRows = pgx.ErrNoRows

func NewDatabasePool(ctx context.Context, logger echo.Logger) (*DatabasePool, error) {
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

	logger.Infof("Database connected %s:%d", config.ConnConfig.Host, config.ConnConfig.Port)

	return &DatabasePool{pool, logger}, nil
}

func (db *DatabasePool) Text(text string) pgtype.Text {
	return pgtype.Text{String: text}
}

func (db *DatabasePool) Acquire(ctx context.Context) (*database.Queries, func(), error) {
	conn, err := db.pool.Acquire(ctx)

	if err != nil {
		return nil, nil, err
	}

	queries := database.New(conn)

	return queries, conn.Release, nil
}

func (db *DatabasePool) WithQueries(ctx context.Context, fn func(*database.Queries) error) error {
	return db.pool.AcquireFunc(ctx, func(c *pgxpool.Conn) error {
		queries := database.New(c)
		return fn(queries)
	})
}

func (db *DatabasePool) Close() {
	db.pool.Close()
	db.logger.Infof("Database disconnected")
}
