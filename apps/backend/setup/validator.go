package setup

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type GoPlaygroundValidator struct {
	validator *validator.Validate
}

func NewValidator() *GoPlaygroundValidator {
	return &GoPlaygroundValidator{validator: validator.New()}
}

func SetupValidator(app *echo.Echo) {
	app.Validator = NewValidator()
}

func (cv *GoPlaygroundValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
