package tsm_auth_domain

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Setup(path string) (error, func(*echo.Echo)) {
	return nil, func(e *echo.Echo) {
		e.POST(path, login)
		e.GET(path, info)
	}
}

func login(c echo.Context) error {
	return c.JSON(http.StatusOK, LoginResult{Token: "", RefreshToken: ""})
}

func info(c echo.Context) error {
	return c.JSON(http.StatusOK, LoginInfo{Data: echo.Map{}})
}
