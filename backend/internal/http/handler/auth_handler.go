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
	UserService *service.UserService
	Logger      *logrus.Logger
}

func NewAuthHandler(logger *logrus.Logger, service *service.UserService) *AuthHandler {
	return &AuthHandler{service, logger}
}

func (handler *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	user := new(model.CreateUserRequest)
	if err := helper.DecodeRequestBody(r, user); err != nil {
		handler.Logger.Warnf("error on decode request body %+v", err)
		helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	result, err := handler.UserService.Register(r.Context(), user)
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
				Message: "user validation error"}, http.StatusUnprocessableEntity)
			return
		default:
			helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusBadRequest)
			return
		}
	}
	helper.WriteJSONResponse(w, &model.Response[*model.TokenResponse]{
		Message: "register success",
		Data:    result,
	}, http.StatusOK)
}

func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	user := new(model.AuthenticateUserRequest)
	if err := helper.DecodeRequestBody(r, user); err != nil {
		helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	result, err := handler.UserService.Login(r.Context(), user)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrUnauthorized):
			helper.WriteJSONResponse(w, &model.ErrResponse{Message: "invalid email or password"}, http.StatusUnauthorized)
			return
		default:
			helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
	}
	helper.WriteJSONResponse(w, &model.Response[*model.TokenResponse]{
		Message: "login success",
		Data:    result,
	}, http.StatusOK)
}

func (handler *AuthHandler) UseGoogleAuth(w http.ResponseWriter, r *http.Request) {
	SCHEMA := "Bearer "
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) == 0 {
		helper.WriteJSONResponse(w, &model.ErrResponse{Message: "invalid token"}, http.StatusBadRequest)
		return
	}
	token := authHeader[len(SCHEMA):]
	response, err := handler.UserService.GoogleAuth(r.Context(), token)
	if err != nil {
		helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	helper.WriteJSONResponse(w, &model.Response[*model.TokenResponse]{Message: "authentication success", Data: response}, http.StatusOK)
}

func (handler *AuthHandler) GetMyProfile(w http.ResponseWriter, r *http.Request) {
	auth, err := handler.UserService.GetCurrentAuth(r.Context())
	if err != nil {
		helper.WriteJSONResponse(w, &model.ErrResponse{Message: "unauthorized user"}, http.StatusUnauthorized)
		return
	}
	helper.WriteJSONResponse(w, &model.Response[*model.AuthResponse]{
		Message: "get auth data success",
		Data:    auth,
	}, http.StatusOK)
}
