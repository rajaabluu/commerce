package service

import (
	"context"
	"fmt"

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

func (service *UserService) Register(ctx context.Context, request *model.CreateUserRequest) (*model.UserResponse, error) {
	tx := service.Database.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := service.Validator.Struct(request); err != nil {
		return nil, err
	}
	customer := &entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Address:  request.Address,
		Contact:  request.Contact,
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error on encrypting password: %v", err)
	}
	customer.Password = string(hash)
	exist := &entity.User{Email: request.Email}
	service.UserRepository.FindOne(tx, exist)
	if exist.ID != 0 {
		return nil, &model.CustomFieldErr{
			CustomErr: &model.CustomErr{
				Inner:   gorm.ErrDuplicatedKey,
				Message: "email has already taken"},
			Field: "email"}
	}
	if err := service.UserRepository.Create(tx, customer); err != nil {
		return nil, fmt.Errorf("error on inserting customer: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("an error occured on transaction %+v", err)
	}
	response := converter.ToUserResponse(customer)
	token, err := helper.GenerateToken(service.Config, response)
	if err != nil {
		return nil, err
	}
	response.Token = token
	return response, nil
}

func (service *UserService) Login(ctx context.Context, request *model.AuthenticateUserRequest) (*model.UserResponse, error) {
	tx := service.Database
	customer := &entity.User{
		Email: request.Email,
	}
	if err := service.UserRepository.FindOne(tx, customer); err != nil {
		service.Logger.Warnf("error on service.findone: %+v", err)
		return nil, model.ErrUnauthorized
	}
	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password)); err != nil {
		service.Logger.Warnf("error on compare password: %+v", err)
		return nil, model.ErrUnauthorized
	}
	response := converter.ToUserResponse(customer)
	token, err := helper.GenerateToken(service.Config, response)
	if err != nil {
		return nil, err
	}
	response.Token = token
	return response, nil
}

func (service *UserService) GetCurrentAuth(ctx context.Context) (*model.AuthResponse, error) {
	tx := service.Database.WithContext(ctx)
	auth := ctx.Value(model.AuthContextKey).(*model.Auth)
	customer := new(entity.User)
	if err := service.UserRepository.FindById(tx, auth.ID, customer); err != nil {
		return nil, err
	}
	response := &model.AuthResponse{
		ID:      customer.ID,
		Name:    customer.Name,
		Email:   customer.Email,
		Role:    customer.Role,
		Address: customer.Address,
		Contact: customer.Contact,
	}
	return response, nil
}
