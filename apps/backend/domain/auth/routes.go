package auth

import (
	"net/http"
	"tsm/domain/user"
	"tsm/util"

	"github.com/labstack/echo/v4"
)

func Setup(path string) util.SetupRoutes {
	return util.Setup(
		func(e *echo.Echo) {
			e.POST(path, login)
			e.GET(path, info)
		},
	)
}

func login(c echo.Context) error {
	return c.JSON(http.StatusOK, LoginResult{Token: "fake-auth", RefreshToken: ""})
}

func info(c echo.Context) error {
	return c.JSON(http.StatusOK, LoginInfo{Data: user.UserData{Name: "Jane Smith", Email: "jane.smith@email.com"}})
}
