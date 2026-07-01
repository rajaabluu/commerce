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
	Config *config.Config
	Logger *logrus.Logger

	DB        *gorm.DB
	Validator *validator.Validate

	UserRepository *repository.UserRepository
}

func NewUserService(
	config *config.Config,
	logger *logrus.Logger,
	DB *gorm.DB,
	validator *validator.Validate,
	repository *repository.UserRepository) *UserService {

	return &UserService{
		Config:         config,
		Logger:         logger,
		DB:             DB,
		Validator:      validator,
		UserRepository: repository,
	}
}

func (s *UserService) Create(ctx context.Context, req *model.CreateUserRequest) (*model.UserResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()

	defer tx.Rollback()

	if err := s.Validator.Struct(req); err != nil {
		return nil, err
	}

	existingUser, err := s.UserRepository.FindByEmail(tx, req.Email)

	if existingUser.Email != "" {
		return nil, echo.ErrUnprocessableEntity
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		s.Logger.Warnf("failed to hashing password: %+v", err)
		return nil, echo.ErrInternalServerError
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(password),
	}

	if req.Phone != "" {
		user.Phone = &req.Phone
	}

	err = s.UserRepository.Create(tx, user)

	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		s.Logger.Warnf("failed to create user: %+v", err)
		return nil, echo.ErrInternalServerError
	}

	res := &model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}

	token, err := helper.GenerateToken(s.Config, res)

	if err != nil {
		s.Logger.Warnf("failed on generating token: %+v", err)
		return nil, echo.ErrInternalServerError
	}

	res.AccessToken = token

	return res, nil

}

func (s *UserService) Login(ctx context.Context, req *model.AuthenticateUserRequest) (*model.UserResponse, error) {
	tx := s.DB.WithContext(ctx)

	user, err := s.UserRepository.FindByEmail(tx, req.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.Logger.Warnf("user not found")
			return nil, echo.ErrUnauthorized
		} else {
			s.Logger.Warn(err.Error())
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

	token, err := helper.GenerateToken(s.Config, res)

	if err != nil {
		s.Logger.Warnf("error on generating token: %+v", err)
		return nil, echo.ErrInternalServerError
	}

	res.AccessToken = token

	return res, nil

}

func (s *UserService) GetAuthenticatedUser(ctx context.Context, ID uint) (*model.UserProfileResponse, error) {
	tx := s.DB.WithContext(ctx)
	user, err := s.UserRepository.FindById(tx, ID)

	if err != nil {
		s.Logger.Warnf("failed to find user by id: %+v", err)
		return nil, err
	}

	res := &model.UserProfileResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}

	if user.Phone != nil {
		res.Phone = user.Phone
	}

	return res, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, req *model.UpdateUserProfileRequest) (*model.UserProfileResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user, err := s.UserRepository.FindById(tx, req.ID)
	if err != nil {
		s.Logger.Warnf("failed on updating user profile: %+v", err)
		return nil, echo.ErrInternalServerError
	}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}

	if req.Phone != nil {
		user.Phone = req.Phone
	}

	err = s.UserRepository.Update(tx, user)

	if err != nil {
		s.Logger.Warnf("failed on updating user profile: %+v", err)
		return nil, echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Logger.Warnf("failed on commit transaction: %+v", err)
		return nil, echo.ErrInternalServerError
	}

	return &model.UserProfileResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Role:  string(user.Role),
	}, nil
}
