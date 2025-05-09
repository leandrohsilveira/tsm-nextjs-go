package main

import (
	"tsm/domain/auth"
	"tsm/util"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	setup := Setup(
		auth.Setup("/auth"),
	)

	if err := setup(e); err != nil {
		e.Logger.Fatal(err)
	}

	if err := e.Start(":4000"); err != nil {
		e.Logger.Fatal(err)
	}
}

func Setup(hooks ...util.SetupRoutes) func(e *echo.Echo) error {
	return func(e *echo.Echo) error {
		for _, hook := range hooks {
			if hook.Err != nil {
				return hook.Err
			}
			hook.Hook(e)
		}
		return nil
	}
}
