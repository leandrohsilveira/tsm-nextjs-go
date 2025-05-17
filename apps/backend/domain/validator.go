package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type FieldError struct {
	Name  string `json:"name"`
	Err   string `json:"error"`
	Value any    `json:"value"`
}

type ValidationData struct {
	Message string       `json:"message"`
	Fields  []FieldError `json:"fields"`
}

type ValidationError struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Fields  []FieldError `json:"fields"`
}

var val = validator.New()

func NewValidationError(result *ValidationData) *ValidationError {
	return &ValidationError{
		Code:    fiber.StatusBadRequest,
		Message: fiber.ErrBadRequest.Message,
		Fields:  result.Fields,
	}
}

func Validate(i any) (*ValidationData, error) {
	err := val.Struct(i)

	errs, ok := err.(validator.ValidationErrors)

	if ok {
		fields := []FieldError{}

		for _, field := range errs {
			fields = append(fields, FieldError{
				Name:  field.Field(),
				Err:   field.Tag(),
				Value: field.Value(),
			})
		}

		data := &ValidationData{
			Message: err.Error(),
			Fields:  fields,
		}
		return data, nil
	}

	if err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return nil, err
	}
	return nil, nil
}

func (err *ValidationError) Error() string {
	return err.Message
}
