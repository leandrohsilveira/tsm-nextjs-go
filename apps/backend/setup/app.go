package setup

import (
	"net/http"
	"tsm/domain"

	"github.com/gofiber/fiber/v2"
)

func SetupApp() *fiber.App {
	return fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})
}

func errorHandler(c *fiber.Ctx, err error) error {
	valErr, ok := err.(*domain.ValidationError)

	if ok {
		return c.Status(valErr.Code).JSON(valErr)
	}

	fiberErr, ok := err.(*fiber.Error)
	if ok {
		return c.Status(fiberErr.Code).JSON(fiberErr)
	}

	return c.Status(http.StatusInternalServerError).JSON(fiber.Error{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	})
}
