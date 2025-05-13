package main

import (
	"context"
	"net/http"
	"tsm/domain"
	"tsm/domain/auth"
	"tsm/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Validator = domain.NewValidator()

	pool, err := domain.NewDatabasePool(context.Background())
	if err != nil {
		app.Logger.Fatal(err)
	}

	defer pool.Close()
	defer app.Logger.Printf("Database disconnected")

	config := pool.Config().ConnConfig
	app.Logger.Printf("Database connected %s:%d", config.Host, config.Port)

	err = Seed(context.Background(), pool)
	if err != nil {
		app.Logger.Fatal(err)
	}

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
