package service

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rajaabluu/commerce/backend/internal/config"
	"github.com/rajaabluu/commerce/backend/internal/entity"
	"github.com/rajaabluu/commerce/backend/internal/helper"
	"github.com/rajaabluu/commerce/backend/internal/model"
	"github.com/rajaabluu/commerce/backend/internal/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	Config         *config.Config
	DB             *gorm.DB
	UserRepository *repository.UserRepository
	Validator      *validator.Validate
	Logger         *logrus.Logger
}

func NewUserService(config *config.Config, validator *validator.Validate, logger *logrus.Logger, DB *gorm.DB, repository *repository.UserRepository) *UserService {
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

	res := &model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}

	token, err := helper.GenerateToken(service.Config, res)

	if err != nil {
		service.Logger.Warnf("failed on generating token: %+v", err)
		return nil, echo.ErrInternalServerError
	}

	res.AccessToken = token

	return res, nil

}

func (service *UserService) Login(ctx context.Context, req *model.AuthenticateUserRequest) (*model.UserResponse, error) {
	tx := service.DB.WithContext(ctx)
	defer tx.Rollback()

	user := new(entity.User)

	err := service.UserRepository.FindByEmail(tx, req.Email, user)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			service.Logger.Warnf("user not found")
			return nil, echo.ErrUnauthorized
		} else {
			service.Logger.Warn(err.Error())
			return nil, err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, echo.ErrUnauthorized
	}

	res := &model.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  string(user.Role),
		Name:  user.Name,
	}

	token, err := helper.GenerateToken(service.Config, res)

	if err != nil {
		service.Logger.Warnf("error on generating token: %+v", err)
		return nil, echo.ErrInternalServerError
	}

	res.AccessToken = token

	return res, nil

}

func (service *UserService) GetAuthenticatedUser(ctx context.Context, ID uint) (*model.AuthenticatedUserResponse, error) {
	tx := service.DB.WithContext(ctx)
	user := new(entity.User)
	if err := service.UserRepository.FindById(tx, ID, user); err != nil {
		service.Logger.Warnf("failed to find user by id: %+v", err)
		return nil, err
	}
	res := &model.AuthenticatedUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}

	if user.Contact != nil {
		res.Contact = user.Contact
	}

	if user.Address != nil {
		res.Address = user.Address
	}

	return res, nil
}
