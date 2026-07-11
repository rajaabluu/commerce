package service

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rajaabluu/commerce/backend/internal/config"
	"github.com/rajaabluu/commerce/backend/internal/entity"
	"github.com/rajaabluu/commerce/backend/internal/helper/mapper"
	"github.com/rajaabluu/commerce/backend/internal/model"
	"github.com/rajaabluu/commerce/backend/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductService struct {
	Config *config.Config
	Logger *logrus.Logger

	DB        *gorm.DB
	Validator *validator.Validate
	Uploader  *cloudinary.Cloudinary

	ProductRepository      *repository.ProductRepository
	ProductImageRepository *repository.ProductImageRepository
}

func NewProductService(
	config *config.Config,
	logger *logrus.Logger,
	DB *gorm.DB,
	validator *validator.Validate,
	uploader *cloudinary.Cloudinary,

	productRepository *repository.ProductRepository,
	productImageRepository *repository.ProductImageRepository,
) *ProductService {

	return &ProductService{
		Config:            config,
		Logger:            logger,
		DB:                DB,
		Validator:         validator,
		ProductRepository: productRepository,
	}
}

func (s *ProductService) Find(ctx context.Context, req *model.GetProductsRequest) ([]*model.ProductResponse, error) {
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
			res = append(res, mapper.ToProductResponse(product))
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

	res := mapper.ToProductResponse(product)

	return res, nil
}

func (s *ProductService) FindByID(ctx context.Context, id uint) (*model.ProductResponse, error) {
	db := s.DB.WithContext(ctx).Preload("Categories").Preload("Images")
	product, err := s.ProductRepository.FindById(db, id)

	if err != nil {
		return nil, err
	}

	res := mapper.ToProductResponse(product)

	return res, nil

}

func (s *ProductService) Update(ctx context.Context, req *model.UpdateProductRequest, ID uint) (*model.ProductResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	product := new(entity.Product)
	product.ID = ID

	if req.Name != nil {
		product.Name = *req.Name
	}

	if req.Description != nil {
		product.Description = *req.Description
	}

	if req.Price != nil {
		product.Price = *req.Price
	}

	if req.Stock != nil {
		product.Stock = *req.Stock
	}

	if req.CategoryIds != nil && len(*req.CategoryIds) > 0 {
		categories := make([]entity.Category, 0, len(*req.CategoryIds))
		for _, id := range *req.CategoryIds {
			categories = append(categories, entity.Category{
				ID: uint(id),
			})
		}

		product := new(entity.Product)
		product.ID = ID

		if err := tx.Model(product).Association("Categories").Replace(categories); err != nil {
			s.Logger.Warnf("error on updating categories association: %+v", err)
			return nil, err
		}
	}

	err := s.ProductRepository.Update(tx, product)
	if err != nil {
		s.Logger.Warnf("error on updating product: %+v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		s.Logger.Warnf("error on commit transaction: %+v", err)
		return nil, err
	}

	res := mapper.ToProductResponse(product)

	return res, nil
}

func (s *ProductService) Delete(ctx context.Context, productID uint) error {
	tx := s.DB.WithContext(ctx).Begin()

	defer tx.Rollback()
	product, err := s.ProductRepository.FindById(tx.Preload("Images"), productID)

	ids := make([]uint, 0, len(product.Images))
	imagePublicIds := make([]string, 0, len(product.Images))

	if len(product.Images) > 0 {
		for _, image := range product.Images {
			ids = append(ids, image.ID)
			imagePublicIds = append(imagePublicIds, image.PublicID)
		}
		_, err := s.Uploader.Admin.DeleteAssets(ctx, admin.DeleteAssetsParams{
			PublicIDs: imagePublicIds,
		})

		if err != nil {
			return err
		}

		tx.Delete(&entity.ProductImage{}, ids)
	}

	if err != nil {
		return err
	}

	if err := s.ProductRepository.Delete(tx, product); err != nil {
		s.Logger.Warnf("error on deleting item: %+v", err)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		s.Logger.Warnf("error on commit transaction: %+v", err)
		return err
	}

	return nil
}
