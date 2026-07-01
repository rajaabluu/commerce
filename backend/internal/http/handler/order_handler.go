package handler

import (
	"errors"

	"github.com/labstack/echo/v4"
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
	return errors.New("not implemented yet")
}
