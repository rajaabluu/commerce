package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/rajaabluu/commerce/backend/internal/config"
	"github.com/rajaabluu/commerce/backend/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductService struct {
	DB                *gorm.DB
	Config            *config.Config
	Logger            *logrus.Logger
	Validator         *validator.Validate
	ProductRepository *repository.ProductRepository
}

func NewProductService(config *config.Config, validator *validator.Validate, logger *logrus.Logger, DB *gorm.DB, repository *repository.ProductRepository) *ProductService {
	return &ProductService{
		Config:            config,
		DB:                DB,
		Validator:         validator,
		Logger:            logger,
		ProductRepository: repository,
	}
}
