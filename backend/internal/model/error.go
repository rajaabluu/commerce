package model

import (
	"errors"
)

type ValidationErr struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var (
	ErrUnauthorized = errors.New("unauthorized user")
)

type CustomErr struct {
	Inner   error
	Message string
}

type CustomFieldErr struct {
	*CustomErr
	Field string
}

func (e *CustomErr) Error() string {
	return e.Message
}

func (e *CustomErr) Unwrap() error {
	return e.Inner
}
