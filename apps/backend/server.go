package main

import (
	"context"
	"net/http"
	"os"
	"tsm/domain"
	"tsm/domain/auth"
	"tsm/util"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Validator = domain.NewValidator()

	connString, isSet := os.LookupEnv("DATABASE_URL")
	if !isSet {
		connString = "postgres://app:password@localhost:5432/app?sslmode=disable"
	}

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		app.Logger.Fatal(err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		app.Logger.Fatal(err)
	}

	defer pool.Close()
	defer app.Logger.Printf("Database disconnected")

	app.Logger.Printf("Database connected %s:%d", config.ConnConfig.Host, config.ConnConfig.Port)

	setup := Setup(
		auth.Setup("/auth", pool),
	)

	if err := setup(app); err != nil {
		app.Logger.Fatal(err)
	}

	if err := app.Start(":4000"); err != nil && err != http.ErrServerClosed {
		app.Logger.Fatal(err)
	}
}

func Setup(hooks ...util.SetupRoutes) func(*echo.Echo) error {
	return func(app *echo.Echo) error {
		for _, hook := range hooks {
			if hook.Err != nil {
				return hook.Err
			}
			hook.Hook(app)
		}
		return nil
	}
}
