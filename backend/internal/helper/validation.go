package helper

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/rajaabluu/ershop-api/internal/model"
)

func CreateValidationErrors(ve validator.ValidationErrors) *[]model.ValidationErr {
	result := make([]model.ValidationErr, len(ve))
	for i, field := range ve {
		result[i] = model.ValidationErr{
			Field:   strings.ToLower(field.Field()),
			Message: GetValidationMessage(field),
		}
	}
	return &result
}

func GetValidationMessage(f validator.FieldError) string {
	switch f.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", strings.ToLower(f.Field()))
	case "email":
		return fmt.Sprintf("%s must be valid email format", strings.ToLower(f.Field()))
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", f.Field(), f.Param())
	}
	return f.Error()
}
