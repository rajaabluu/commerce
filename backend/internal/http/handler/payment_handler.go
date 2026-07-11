package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rajaabluu/commerce/backend/internal/model"
	"github.com/rajaabluu/commerce/backend/internal/service"
	"github.com/sirupsen/logrus"
)

type PaymentHandler struct {
	Logger *logrus.Logger

	PaymentService *service.PaymentService
}

func NewPaymentHandler(logger *logrus.Logger, paymentService *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		Logger:         logger,
		PaymentService: paymentService,
	}
}

func HandlePaymentNotification(c echo.Context) error {
	req := new(model.MidtransNotificationRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, &model.ErrorResponse{
			Message: "invalid notification req",
		})
	}
	return errors.New("not implemented yet")
}
