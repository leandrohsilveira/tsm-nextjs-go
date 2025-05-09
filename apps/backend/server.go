package main

import (
	"net/http"
	"tsm/domain/auth"
	"tsm/util"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	setup := Setup(
		auth.Setup("/auth"),
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
