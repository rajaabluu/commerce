package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rajaabluu/commerce/backend/internal/config"
	"github.com/rajaabluu/commerce/backend/internal/entity"
	"github.com/rajaabluu/commerce/backend/internal/model"
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

func (service *ProductService) Create(ctx context.Context, req *model.CreateProductRequest) (*model.ProductResponse, error) {
	tx := service.DB.WithContext(ctx).Begin()

	defer tx.Rollback()

	if err := service.Validator.Struct(req); err != nil {
		return nil, err
	}

	ids := req.CategoryIds

	var categories []entity.Category

	result := tx.Where("id IN ?", ids).Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	product := &entity.Product{
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
		Categories:  categories,
	}

	if err := service.ProductRepository.Create(tx, product); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		service.Logger.Warnf("failed to create user: %+v", err)
		return nil, echo.ErrInternalServerError
	}

	res := &model.ProductResponse{
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
		Price:       product.Price,
		Categories:  &product.Categories,
	}

	return res, nil
}
