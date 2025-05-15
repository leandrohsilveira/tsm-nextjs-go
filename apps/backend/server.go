package main

import (
	"context"
	"net/http"
	"tsm/domain"
	"tsm/domain/auth"
	"tsm/setup"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	setup.SetupLogger(app)
	setup.SetupValidator(app)

	pool, err := domain.NewDatabasePool(context.Background(), app.Logger)

	if err != nil {
		app.Logger.Fatal(err)
	}

	defer pool.Close()

	if err = Seed(context.Background(), app.Logger, pool); err != nil {
		app.Logger.Fatal(err)
	}

	setupRoutes := setup.Routes(
		auth.Routes("/auth", pool),
	)

	if err := setupRoutes(app); err != nil {
		app.Logger.Fatal(err)
	}

	if err := app.Start(":4000"); err != nil && err != http.ErrServerClosed {
		app.Logger.Fatal(err)
	}
}
