package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rajaabluu/ershop-api/internal/entity"
	"github.com/rajaabluu/ershop-api/internal/helper"
	"github.com/rajaabluu/ershop-api/internal/helper/converter"
	"github.com/rajaabluu/ershop-api/internal/model"
	"github.com/rajaabluu/ershop-api/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	Config         *viper.Viper
	Validator      *validator.Validate
	Database       *gorm.DB
	Logger         *logrus.Logger
	UserRepository *repository.UserRepository
}

func NewUserService(config *viper.Viper, database *gorm.DB, validator *validator.Validate, logger *logrus.Logger, repository *repository.UserRepository) *UserService {
	return &UserService{
		Config:         config,
		Database:       database,
		Validator:      validator,
		Logger:         logger,
		UserRepository: repository,
	}
}

func (service *UserService) Verify(token string) (*model.Auth, error) {
	claims, err := helper.ValidateToken(service.Config, token)
	if err != nil {
		return nil, err
	}
	id := claims["id"].(float64)
	return &model.Auth{ID: uint(id)}, nil
}

func (service *UserService) Register(ctx context.Context, request *model.CreateUserRequest) (*model.TokenResponse, error) {
	tx := service.Database.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := service.Validator.Struct(request); err != nil {
		return nil, err
	}
	user := &entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Address:  request.Address,
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error on encrypting password: %v", err)
	}
	user.Password = string(hash)
	exist := &entity.User{Email: request.Email}
	service.UserRepository.FindOne(tx, exist)
	if exist.ID != 0 {
		return nil, &model.CustomFieldErr{
			CustomErr: &model.CustomErr{
				Inner:   gorm.ErrDuplicatedKey,
				Message: "email has already taken"},
			Field: "email"}
	}
	if err := service.UserRepository.Create(tx, user); err != nil {
		return nil, fmt.Errorf("error on inserting user: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("an error occured on transaction %+v", err)
	}
	response := converter.ToUserResponse(user)
	token, err := helper.GenerateToken(service.Config, response)
	if err != nil {
		return nil, err
	}
	return &model.TokenResponse{AccessToken: token}, nil
}

func (service *UserService) Login(ctx context.Context, request *model.AuthenticateUserRequest) (*model.TokenResponse, error) {
	tx := service.Database
	user := &entity.User{
		Email: request.Email,
	}
	if err := service.UserRepository.FindOne(tx, user); err != nil {
		service.Logger.Warnf("error on service.findone: %+v", err)
		return nil, model.ErrUnauthorized
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		service.Logger.Warnf("error on compare password: %+v", err)
		return nil, model.ErrUnauthorized
	}
	response := converter.ToUserResponse(user)
	token, err := helper.GenerateToken(service.Config, response)
	if err != nil {
		return nil, err
	}
	return &model.TokenResponse{AccessToken: token}, nil
}

func (service *UserService) GoogleAuth(ctx context.Context, token string) (*model.TokenResponse, error) {
	tx := service.Database.WithContext(ctx).Begin()
	defer tx.Rollback()
	req, err := http.NewRequest(http.MethodGet, "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	googleRes := new(model.GoogleAuthResponse)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error on getting user data: %+v", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error on reading body response from google, %v", err)
	}
	err = json.Unmarshal(body, googleRes)
	if err != nil {
		return nil, fmt.Errorf("error on unmarshal json: %+v", err)
	}
	user := &entity.User{Email: googleRes.Email}
	err = service.UserRepository.FindOne(tx, user)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			user.Name = googleRes.Name
			err := service.UserRepository.Create(tx, user)
			if err != nil {
				return nil, fmt.Errorf("error on create user: %+v", err)
			}
		default:
			return nil, err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("error on commit transaction: %+v", err)
	}
	token, err = helper.GenerateToken(service.Config, &model.UserResponse{ID: user.ID})
	if err != nil {
		return nil, fmt.Errorf("error on generate token, %+v", err)
	}
	return &model.TokenResponse{AccessToken: token}, nil
}

func (service *UserService) GetCurrentAuth(ctx context.Context) (*model.AuthResponse, error) {
	tx := service.Database.WithContext(ctx)
	auth := ctx.Value(model.AuthContextKey).(*model.Auth)
	user := new(entity.User)
	if err := service.UserRepository.FindById(tx, auth.ID, user); err != nil {
		return nil, err
	}
	response := &model.AuthResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Role:    user.Role,
		Address: user.Address,
	}
	return response, nil
}
