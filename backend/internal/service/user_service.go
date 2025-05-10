package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rajaabluu/commerce/backend/internal/entity"
	"github.com/rajaabluu/commerce/backend/internal/model"
	"github.com/rajaabluu/commerce/backend/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	Config         *viper.Viper
	DB             *gorm.DB
	UserRepository *repository.UserRepository
	Validator      *validator.Validate
	Logger         *logrus.Logger
}

func NewUserService(config *viper.Viper, validator *validator.Validate, logger *logrus.Logger, DB *gorm.DB, repository *repository.UserRepository) *UserService {
	return &UserService{
		Config:         config,
		DB:             DB,
		UserRepository: repository,
		Validator:      validator,
		Logger:         logger,
	}
}

func (service *UserService) Create(ctx context.Context, req *model.CreateUserRequest) (*model.UserResponse, error) {
	tx := service.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := service.Validator.Struct(req); err != nil {
		return nil, err
	}

	userExist := new(entity.User)

	service.UserRepository.FindByEmail(tx, req.Email, userExist)

	if userExist.Email != "" {
		return nil, echo.ErrUnprocessableEntity
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		service.Logger.Warnf("failed to hashing password: %+v", err)
		return nil, echo.ErrInternalServerError
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(password),
	}

	if req.Contact != "" {
		user.Contact = &req.Contact
	}

	if req.Address != "" {
		user.Address = &req.Address
	}

	if err := service.UserRepository.Create(tx, user); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		service.Logger.Warnf("failed to create user: %+v", err)
		return nil, echo.ErrInternalServerError
	}

	return &model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}, nil

}
