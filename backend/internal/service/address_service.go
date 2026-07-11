package service

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/rajaabluu/commerce/backend/internal/config"
	"github.com/rajaabluu/commerce/backend/internal/helper/mapper"
	"github.com/rajaabluu/commerce/backend/internal/model"
	"github.com/rajaabluu/commerce/backend/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AddressService struct {
	Config *config.Config
	Logger *logrus.Logger

	DB        *gorm.DB
	Validator *validator.Validate

	AddressRepository *repository.AddressRepository
}

func NewAddressService(

	config *config.Config,
	logger *logrus.Logger,

	db *gorm.DB,
	validator *validator.Validate,

	addressRepository *repository.AddressRepository,
) *AddressService {
	return &AddressService{
		Config:            config,
		Logger:            logger,
		DB:                db,
		Validator:         validator,
		AddressRepository: addressRepository,
	}
}

func (s *AddressService) Create(ctx context.Context, userID uint, req *model.CreateAddressRequest) (*model.AddressResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()

	defer tx.Rollback()

	address := mapper.ToAddressEntity(req, userID)

	if err := s.AddressRepository.Create(tx, address); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return nil, errors.New("not implemented yet")
}
