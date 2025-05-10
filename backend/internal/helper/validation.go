package helper

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/rajaabluu/commerce/backend/internal/model"
)

func GenerateValidationError(ve validator.ValidationErrors) *[]model.ValidationErr {
	result := make([]model.ValidationErr, len(ve))
	for i, e := range ve {
		result[i] = model.ValidationErr{
			Field:   strings.ToLower(e.Field()),
			Message: GetValidationMessage(e),
		}
	}
	return &result
}

func GetValidationMessage(field validator.FieldError) string {
	switch field.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", strings.ToLower(field.Field()))
	case "email":
		return fmt.Sprintf("%s must be valid email format", strings.ToLower(field.Field()))
	case "min":
		return fmt.Sprintf("%s must be atleast %s", strings.ToLower(field.Field()), field.Param())
	}
	return field.Error()
}
