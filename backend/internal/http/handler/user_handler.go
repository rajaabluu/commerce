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
	userRequest := new(model.CreateUserRequest)
	if err := c.Bind(userRequest); err != nil {
		handler.Logger.Warnf("error on decoding body request: %+v", err)
		return err
	}

	userResponse, err := handler.UserService.Create(c.Request().Context(), userRequest)

	if err != nil {
		var ve validator.ValidationErrors
		switch {

		case errors.As(err, &ve):
			return c.JSON(http.StatusUnprocessableEntity, &model.ErrorResponse{
				Message: "validation error",
				Error:   helper.GenerateValidationError(ve),
			})

		case errors.Is(err, echo.ErrUnprocessableEntity):
			return c.JSON(http.StatusUnprocessableEntity, &model.ErrorResponse{
				Message: "validation error",
				Error: []*model.ValidationErr{{
					Field:   "email",
					Message: "email has already used",
				}},
			})
		}

	}

	token, err := helper.GenerateToken(handler.UserService.Config, userResponse)

	if err != nil {
		handler.Logger.Warnf("error on generating token: %+v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, &model.Response[any]{
		Message: "user succesfully registered",
		Data: &model.UserTokenResponse{
			AccessToken: token,
		},
	})

}
