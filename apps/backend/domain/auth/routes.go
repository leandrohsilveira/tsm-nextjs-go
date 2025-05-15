package auth

import (
	"net/http"
	"tsm/domain"
	"tsm/domain/user"
	"tsm/setup"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AuthRoutes struct {
	pool domain.DatabasePool
}

func Routes(path string, pool domain.DatabasePool) setup.SetupRoutesResult {

	routes := AuthRoutes{pool}

	return setup.SetupRoutes(
		func(e *echo.Echo) {
			e.POST(path, routes.login)
			e.GET(path, routes.info)
		},
	)
}

func (routes *AuthRoutes) login(c echo.Context) error {
	payload := new(LoginPayload)

	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	service := NewService(user.NewService(routes.pool))

	data, err := service.Login(c.Request().Context(), *payload)

	if err == user.ErrIncorrectUsernamePassword {
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (routes *AuthRoutes) info(c echo.Context) error {
	authorization := c.Request().Header.Get("authorization")
	if authorization == "" {
		return echo.ErrUnauthorized
	}

	// TODO: validate and decode bearer token instead of inferring it as user id
	userId, err := uuid.Parse(authorization)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	service := user.NewService(routes.pool)

	data, err := service.GetById(c.Request().Context(), userId)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, LoginInfo{Data: *data})
}
