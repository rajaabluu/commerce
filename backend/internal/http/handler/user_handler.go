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
	Logger      *logrus.Logger
	UserService *service.UserService
}

func NewUserHandler(logger *logrus.Logger, service *service.UserService) *UserHandler {
	return &UserHandler{
		Logger:      logger,
		UserService: service,
	}
}

func (h *UserHandler) Register(c echo.Context) error {

	userRequest := new(model.CreateUserRequest)

	if err := c.Bind(userRequest); err != nil {
		h.Logger.Warnf("error on decoding body request: %+v", err)
		return err
	}

	userResponse, err := h.UserService.Create(c.Request().Context(), userRequest)

	if err != nil {
		var ve validator.ValidationErrors
		switch {
		case errors.As(err, &ve):
			return c.JSON(http.StatusUnprocessableEntity, &model.ErrorResponse{
				Message: "validation error",
				Errors:  helper.GenerateValidationError(ve),
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
	return c.JSON(http.StatusOK, &model.Response[any]{
		Message: "user succesfully registered",
		Data:    userResponse,
	})

}

func (h *UserHandler) Login(c echo.Context) error {
	req := new(model.AuthenticateUserRequest)
	if err := c.Bind(req); err != nil {
		h.Logger.Warnf("error on parsing body: %+v", err)
		return c.JSON(http.StatusBadRequest, &model.ErrorResponse{
			Message: err.Error(),
		})
	}
	res, err := h.UserService.Login(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &model.ErrorResponse{
			Message: "incorrect email or password",
		})
	}
	return c.JSON(http.StatusOK, &model.Response[*model.UserResponse]{
		Message: "login success",
		Data:    res,
	})
}

func (h *UserHandler) GetAuthenticatedUser(c echo.Context) error {
	ID := uint(c.Get("userId").(float64))
	res, err := h.UserService.GetAuthenticatedUser(c.Request().Context(), ID)
	if err != nil {
		h.Logger.Warnf("failed to get authenticated user: %+v", err)
		return c.JSON(http.StatusUnauthorized, &model.ErrorResponse{
			Message: "unauthorized user",
		})
	}
	return c.JSON(http.StatusOK, &model.Response[*model.UserResponse]{
		Message: "data sucsefully retrieved",
		Data:    res,
	})
}

func (h *UserHandler) UpdateProfile(c echo.Context) error {
	req := new(model.UpdateUserProfileRequest)
	req.ID = uint(c.Get("userId").(float64))
	if err := c.Bind(req); err != nil {
		h.Logger.Warnf("error on parsing body: %+v", err)
		return c.JSON(http.StatusBadRequest, &model.ErrorResponse{Message: err.Error()})
	}
	res, err := h.UserService.UpdateProfile(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.ErrorResponse{Message: "failed to update user data"})
	}

	return c.JSON(http.StatusOK, &model.Response[*model.UserResponse]{
		Message: "success update profile",
		Data:    res,
	})
}
