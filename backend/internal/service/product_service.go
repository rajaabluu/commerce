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

func (s *ProductService) GetProducts(ctx context.Context, req *model.GetProductsRequest) ([]*model.ProductResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()

	defer tx.Rollback()

	filter := &model.ProductFilter{
		Categories: req.Categories,
		SortBy:     req.SortBy,
		SortOrder:  req.SortOrder,
		Limit:      req.Limit,
		Offset:     (req.Page - 1) * req.Limit,
	}

	if req.Search != "" {
		filter.Search = req.Search
	}

	if req.MinPrice > 0 {
		filter.MinPrice = req.MinPrice
	}

	if req.MaxPrice > 0 {
		filter.MaxPrice = req.MaxPrice
	}

	res := make([]*model.ProductResponse, 0)

	products, err := s.ProductRepository.Find(tx, filter)

	if err != nil {
		s.Logger.Errorf("error on getting products: %+v", err.Error())
	}

	if len(products) > 0 {
		for _, product := range products {
			var categories []*model.Category
			for _, category := range product.Categories {
				categories = append(categories, &model.Category{
					ID:   category.ID,
					Name: category.Name,
				})
			}
			res = append(res, &model.ProductResponse{
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				Stock:       product.Stock,
				Categories:  categories,
			})
		}
	}

	return res, nil

}

func (s *ProductService) Create(ctx context.Context, req *model.CreateProductRequest) (*model.ProductResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()

	defer tx.Rollback()

	if err := s.Validator.Struct(req); err != nil {
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

	err := s.ProductRepository.Create(tx, product)

	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		s.Logger.Warnf("failed to create user: %+v", err)
		return nil, echo.ErrInternalServerError
	}

	var pCategories []*model.Category

	for _, category := range product.Categories {
		pCategories = append(pCategories, &model.Category{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	res := &model.ProductResponse{
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
		Price:       product.Price,
		Categories:  pCategories,
	}

	return res, nil
}
