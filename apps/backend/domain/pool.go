package domain

import (
	"context"
	"os"
	"tsm/database"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type DatabasePool interface {
	Text(string) pgtype.Text
	Acquire(context.Context) (*database.Queries, func(), error)
	WithQueries(context.Context, func(*database.Queries) error) error
	Close()
}

type databasePool struct {
	pool *pgxpool.Pool
}

var ErrNoRows = pgx.ErrNoRows

func NewDatabasePool(ctx context.Context) (DatabasePool, error) {
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

	log.Info().
		Str("host", config.ConnConfig.Host).
		Uint16("port", config.ConnConfig.Port).
		Msg("Database connected")

	return &databasePool{pool}, nil
}

func (db *databasePool) Text(text string) pgtype.Text {
	return pgtype.Text{String: text, Valid: true}
}

func (db *databasePool) Acquire(ctx context.Context) (*database.Queries, func(), error) {
	conn, err := db.pool.Acquire(ctx)

	if err != nil {
		return nil, nil, err
	}

	queries := database.New(conn)

	return queries, conn.Release, nil
}

func (db *databasePool) WithQueries(ctx context.Context, fn func(*database.Queries) error) error {
	return db.pool.AcquireFunc(ctx, func(c *pgxpool.Conn) error {
		queries := database.New(c)
		return fn(queries)
	})
}

func (db *databasePool) Close() {
	db.pool.Close()
	log.Info().Msg("Database disconnected")
}
