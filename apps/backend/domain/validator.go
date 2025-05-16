package domain

import (
	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Name  string `json:"name"`
	Err   string `json:"error"`
	Value any    `json:"value"`
}

type ValidationResult struct {
	Validated bool
	Err       error        `json:"message"`
	Fields    []FieldError `json:"fields"`
}

var val = validator.New()

func Validate(i any) ValidationResult {
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

		return ValidationResult{
			Err:       err,
			Validated: true,
			Fields:    fields,
		}
	}

	if err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return ValidationResult{Err: err, Validated: false}
	}
	return ValidationResult{Validated: true}
}
