package main

import (
	"tsm/domain/auth"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	err, auth := tsm_auth_domain.Setup("/auth")

	if err != nil {
		e.Logger.Fatal(err)
	}

	auth(e)

	e.Logger.Fatal(e.Start(":4000"))
}
