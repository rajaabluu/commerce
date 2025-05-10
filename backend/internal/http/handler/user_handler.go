package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rajaabluu/commerce/backend/internal/model"
	"github.com/rajaabluu/commerce/backend/internal/service"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserService *service.UserService
	Logger      *logrus.Logger
}

func NewUserHandler(logger *logrus.Logger, service *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
		Logger:      logger,
	}
}

func (handler *UserHandler) Register(c echo.Context) error {
	userRequest := new(*model.CreateUserRequest)
	if err := c.Bind(userRequest); err != nil {
		handler.Logger.Warnf("error on decoding body request: %+v", err)
		return err
	}
	return c.JSON(http.StatusOK, userRequest)
	// userResponse, err := handler.UserService.Create(c.Request().Context(), *userRequest)
}
