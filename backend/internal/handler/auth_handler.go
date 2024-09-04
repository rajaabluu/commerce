package handler

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rajaabluu/ershop-api/internal/helper"
	"github.com/rajaabluu/ershop-api/internal/model"
	"github.com/rajaabluu/ershop-api/internal/service"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	CustomerService *service.CustomerService
	Logger          *logrus.Logger
}

func NewAuthHandler(logger *logrus.Logger, service *service.CustomerService) *AuthHandler {
	return &AuthHandler{service, logger}
}

func (handler *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	customer := new(model.CreateCustomerRequest)
	if err := helper.DecodeRequestBody(r, customer); err != nil {
		handler.Logger.Warnf("error on decode request body %+v", err)
		helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	result, err := handler.CustomerService.Register(r.Context(), customer)
	if err != nil {
		var fe *model.CustomFieldErr
		var ve validator.ValidationErrors
		switch {
		case errors.As(err, &fe):
			helper.WriteJSONResponse(w,
				&model.ErrResponse{
					Message: fe.Inner.Error(),
					Errors: []*model.ValidationErr{{
						Field:   fe.Field,
						Message: fe.Error(),
					}}},
				http.StatusUnprocessableEntity)
			return
		case errors.As(err, &ve):
			helper.WriteJSONResponse(w, &model.ErrResponse{
				Errors:  helper.CreateValidationErrors(ve),
				Message: "customer validation error"}, http.StatusUnprocessableEntity)
			return
		default:
			helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusBadRequest)
			return
		}
	}
	token, err := helper.GenerateToken(handler.CustomerService.Config, result)
	if err != nil {
		helper.WriteJSONResponse(w, map[string]any{"errors": err.Error()}, http.StatusInternalServerError)
	}
	helper.WriteJSONResponse(w, &model.ResponseWithToken[*model.CustomerResponse]{
		Response: &model.Response[*model.CustomerResponse]{
			Message: "register successful",
			Data:    result,
		},
		Token: token,
	}, http.StatusOK)
}

func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	customer := new(model.AuthenticateCustomerRequest)
	if err := helper.DecodeRequestBody(r, customer); err != nil {
		helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	result, err := handler.CustomerService.Login(r.Context(), customer)
	if err != nil {
		helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusUnauthorized)
		return
	}
	token, err := helper.GenerateToken(handler.CustomerService.Config, result)
	if err != nil {
		handler.Logger.Warnf("error on generating token: %+v", err)
		return
	}
	helper.WriteJSONResponse(w, &model.ResponseWithToken[*model.CustomerResponse]{
		Response: &model.Response[*model.CustomerResponse]{
			Message: "Login success",
			Data:    result,
		},
		Token: token,
	}, http.StatusOK)
}
