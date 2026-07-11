package handler

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rajaabluu/commerce/backend/internal/helper"
	"github.com/rajaabluu/commerce/backend/internal/model"
	"github.com/rajaabluu/commerce/backend/internal/service"
	"github.com/sirupsen/logrus"
)

type AddressHandler struct {
	Logger *logrus.Logger

	AddressService *service.AddressService
}

func NewAddressHandler(logger *logrus.Logger, addressService *service.AddressService) *AddressHandler {
	return &AddressHandler{
		Logger:         logger,
		AddressService: addressService,
	}
}

func (h *AddressHandler) CreateNewAddress(c echo.Context) error {
	req := new(model.CreateAddressRequest)

	if err := c.Bind(req); err != nil {
		var ve validator.ValidationErrors
		switch {
		case errors.As(err, &ve):
			errors := helper.GenerateValidationError(ve)
			return c.JSON(http.StatusUnprocessableEntity, &model.ErrorResponse{
				Message: "validation errors",
				Error:   errors,
			})
		}
	}

	return nil
}
