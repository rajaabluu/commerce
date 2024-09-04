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

type CustomerService struct {
	Config             *viper.Viper
	Validator          *validator.Validate
	Database           *gorm.DB
	Logger             *logrus.Logger
	CustomerRepository *repository.CustomerRepository
}

func NewCustomerService(config *viper.Viper, database *gorm.DB, validator *validator.Validate, logger *logrus.Logger, repository *repository.CustomerRepository) *CustomerService {
	return &CustomerService{
		Config:             config,
		Database:           database,
		Validator:          validator,
		Logger:             logger,
		CustomerRepository: repository,
	}
}

func (service *CustomerService) Verify(token string) (*model.Auth, error) {
	claims, err := helper.ValidateToken(service.Config, token)
	if err != nil {
		return nil, err
	}
	id := claims["id"].(float64)
	return &model.Auth{ID: uint(id)}, nil
}

func (service *CustomerService) Register(ctx context.Context, request *model.CreateCustomerRequest) (*model.CustomerResponse, error) {
	tx := service.Database.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := service.Validator.Struct(request); err != nil {
		return nil, err
	}
	customer := &entity.Customer{
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
	exist := &entity.Customer{Email: request.Email}
	service.CustomerRepository.FindOne(tx, exist)
	if exist.ID != 0 {
		return nil, &model.CustomFieldErr{
			CustomErr: &model.CustomErr{
				Inner:   gorm.ErrDuplicatedKey,
				Message: "email has already taken"},
			Field: "email"}
	}
	if err := service.CustomerRepository.Create(tx, customer); err != nil {
		return nil, fmt.Errorf("error on inserting customer: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("an error occured on transaction %+v", err)
	}
	return converter.ToCustomerResponse(customer), nil
}

func (service *CustomerService) Login(ctx context.Context, request *model.AuthenticateCustomerRequest) (*model.CustomerResponse, error) {
	tx := service.Database
	customer := &entity.Customer{
		Email: request.Email,
	}
	if err := service.CustomerRepository.FindOne(tx, customer); err != nil {
		service.Logger.Warnf("error on service.findone: %+v", err)
		return nil, model.ErrUnauthorized
	}
	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password)); err != nil {
		service.Logger.Warnf("error on compare password: %+v", err)
		return nil, model.ErrUnauthorized
	}
	return converter.ToCustomerResponse(customer), nil
}
