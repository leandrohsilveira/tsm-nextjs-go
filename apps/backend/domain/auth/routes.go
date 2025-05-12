package auth

import (
	"net/http"
	"tsm/domain/user"
	"tsm/util"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type AuthRoutes struct {
	pool *pgxpool.Pool
}

func Setup(path string, pool *pgxpool.Pool) util.SetupRoutes {

	routes := AuthRoutes{pool}

	return util.Setup(
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

	service := user.NewService(routes.pool)

	user, err := service.GetByEmailAndPassword(c.Request().Context(), payload.Username, payload.Password)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, LoginResult{Token: user.ID, RefreshToken: ""})
}

func (routes *AuthRoutes) info(c echo.Context) error {
	return c.JSON(http.StatusOK, LoginInfo{Data: user.UserData{Name: "Jane Smith", Email: "jane.smith@email.com"}})
}
