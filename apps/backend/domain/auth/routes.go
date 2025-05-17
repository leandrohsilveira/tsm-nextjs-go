package auth

import (
	"net/http"
	"tsm/domain"
	"tsm/domain/user"

	"github.com/gofiber/fiber/v2"
)

type AuthRoutes struct {
	pool domain.DatabasePool
}

func Routes(pool domain.DatabasePool) *fiber.App {
	app := fiber.New()

	routes := AuthRoutes{pool}

	app.Post("/", routes.login).Name("Login route")
	app.Get("/", routes.info).Name("Current user info route")

	return app
}

func (routes *AuthRoutes) login(c *fiber.Ctx) error {
	payload := new(LoginPayload)

	if err := c.BodyParser(payload); err != nil {
		return domain.NewHttpError(http.StatusBadRequest, err)
	}

	result := domain.Validate(payload)

	if result.Validated && result.Err != nil {
		return c.Status(http.StatusBadRequest).JSON(result)
	}

	if result.Err != nil {
		return domain.NewHttpError(http.StatusBadRequest, result.Err)
	}

	service := NewService(user.NewService(routes.pool))

	data, err := service.Login(c.UserContext(), *payload)

	if err == user.ErrIncorrectUsernamePassword {
		return domain.NewHttpError(http.StatusForbidden, err)
	}

	if err != nil {
		return err
	}

	return c.JSON(data)
}

func (routes *AuthRoutes) info(c *fiber.Ctx) error {
	payload := new(LoginInfoPayload)
	if err := c.ReqHeaderParser(payload); err != nil {
		return err
	}

	result := domain.Validate(payload)
	if result.Err != nil && result.Validated {
		return domain.ErrUnauthorized
	}
	if result.Err != nil {
		return result.Err
	}

	service := NewService(user.NewService(routes.pool))

	data, err := service.GetCurrentUser(c.UserContext(), *payload)
	if err != nil {
		return err
	}
	if data == nil {
		return domain.ErrUnauthorized
	}

	return c.JSON(LoginInfo{Data: *data})
}
