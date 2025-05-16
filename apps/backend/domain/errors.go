package domain

import "github.com/gofiber/fiber/v2"

var ErrUnauthorized = fiber.ErrUnauthorized

func NewHttpError(code int, err error) *fiber.Error {
	return &fiber.Error{
		Code:    code,
		Message: err.Error(),
	}
}
