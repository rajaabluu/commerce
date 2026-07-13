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

type OrderHandler struct {
	Logger *logrus.Logger

	OrderService *service.OrderService
}

func NewOrderHandler(logger *logrus.Logger, orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{
		Logger:       logger,
		OrderService: orderService,
	}
}

func (h *OrderHandler) CreateNewOrder(c echo.Context) error {
	req := new(model.CreateOrderRequest)

	if err := c.Bind(req); err != nil {
		h.Logger.Warnf("error on decoding request body: %+v", err)
		return c.JSON(http.StatusBadRequest, &model.ErrorResponse{
			Message: "invalid request format",
		})
	}

	userID := uint(c.Get("userId").(float64))

	res, err := h.OrderService.Create(c.Request().Context(), userID, req)

	if err != nil {
		var ve validator.ValidationErrors
		switch {
		case errors.As(err, &ve):
			errors := helper.GenerateValidationError(ve)
			return c.JSON(http.StatusUnprocessableEntity, &model.ErrorResponse{
				Message: "validation error",
				Errors:  errors,
			})
		default:
			return c.JSON(http.StatusInternalServerError, &model.ErrorResponse{
				Message: "internal server error",
				Error:   err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, &model.Response[*model.OrderResponse]{
		Message: "order successfully created",
		Data:    res,
	})
}
