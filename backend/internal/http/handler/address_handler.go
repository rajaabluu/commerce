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

func (h *AddressHandler) GetAllAddresses(c echo.Context) error {
	userID := uint(c.Get("userId").(float64))

	res, err := h.AddressService.GetAll(c.Request().Context(), userID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.ErrorResponse{
			Message: "unexpected error occuredd",
		})
	}

	return c.JSON(http.StatusOK, &model.Response[[]*model.AddressResponse]{
		Message: "addresses retrieved successfully",
		Data:    res,
	})
}

func (h *AddressHandler) CreateNewAddress(c echo.Context) error {
	req := new(model.CreateAddressRequest)

	if err := c.Bind(req); err != nil {
		h.Logger.Warnf("error on decoding request body: %+v", err)
	}

	userID := uint(c.Get("userId").(float64))

	res, err := h.AddressService.Create(c.Request().Context(), userID, req)

	var ve validator.ValidationErrors
	switch {
	case errors.As(err, &ve):
		errors := helper.GenerateValidationError(ve)
		return c.JSON(http.StatusUnprocessableEntity, &model.ErrorResponse{
			Message: "validation errors",
			Errors:  errors,
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.ErrorResponse{
			Message: "unexpected error occuredd",
		})
	}

	return c.JSON(http.StatusOK, &model.Response[*model.AddressResponse]{
		Message: "address created successfully",
		Data:    res,
	})
}
