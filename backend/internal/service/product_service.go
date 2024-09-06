package service

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/rajaabluu/ershop-api/internal/entity"
	"github.com/rajaabluu/ershop-api/internal/helper/converter"
	"github.com/rajaabluu/ershop-api/internal/model"
	"github.com/rajaabluu/ershop-api/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductService struct {
	Validator         *validator.Validate
	Logger            *logrus.Logger
	Database          *gorm.DB
	ProductRepository *repository.ProductRepository
	Uploader          *cloudinary.Cloudinary
}

func NewProductService(database *gorm.DB, validator *validator.Validate, uploader *cloudinary.Cloudinary, logger *logrus.Logger, repository *repository.ProductRepository) *ProductService {
	return &ProductService{
		Database:          database,
		Validator:         validator,
		Logger:            logger,
		ProductRepository: repository,
		Uploader:          uploader,
	}
}

func (service *ProductService) GetCategories(ctx context.Context) ([]*model.Category, error) {
	tx := service.Database.WithContext(ctx)
	var categories []*entity.Category
	err := tx.Model(new(entity.Category)).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	result := make([]*model.Category, len(categories))
	for i, category := range categories {
		result[i] = &model.Category{ID: category.ID, Name: category.Name}
	}
	return result, nil
}

func (service *ProductService) CreateProduct(ctx context.Context, request *model.CreateProductRequest) (*model.ProductResponse, error) {
	tx := service.Database.Model(&entity.Product{}).Preload("Categories").WithContext(ctx).Begin()
	defer tx.Rollback()
	product := &entity.Product{
		Name:        request.Name,
		Description: request.Description,
		Quantity:    request.Quantity,
		Price:       request.Price,
	}
	categories := make([]*entity.Category, len(request.Categories))
	for i, category := range request.Categories {
		item := new(entity.Category)
		if err := service.Database.Where("name = ?", category).First(item).Error; err != nil {
			service.Logger.Errorf("error on finding categories %+v", err)
			break
		}
		categories[i] = item
	}
	product.Categories = categories
	images := make([]*entity.ProductImage, len(request.Images))
	for i, val := range request.Images {
		file, err := val.Open()
		if err != nil {
			service.Logger.Warnf("error on opening file: %+v", err)
		}
		defer file.Close()
		res, err := service.Uploader.Upload.Upload(ctx, file, uploader.UploadParams{})
		if err != nil {
			service.Logger.Warnf("error on uploading image: %+v", err)
		}
		images[i] = &entity.ProductImage{Source: res.SecureURL, PublicID: res.PublicID}
	}
	product.Images = images
	if err := service.ProductRepository.Create(tx.Preload("Categories"), product); err != nil {
		return nil, fmt.Errorf("error on creating product : %+v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("error on commit transaction: %+v", err)
	}
	return converter.ToProductResponse(product), nil
}

func (service *ProductService) GetAllProducts(ctx context.Context) ([]*model.ProductResponse, error) {
	tx := service.Database.Model(&entity.Product{}).Preload("Categories").Preload("Images").WithContext(ctx)
	var products []*entity.Product
	if err := service.ProductRepository.Find(tx, &entity.Product{}, &products); err != nil {
		return nil, fmt.Errorf("error on find products %+v", err)
	}
	responses := make([]*model.ProductResponse, len(products))
	for i, product := range products {
		item := converter.ToProductResponse(product)
		responses[i] = item
	}
	return responses, nil
}

func (service *ProductService) GetProductsByCategory(ctx context.Context, categories []string) ([]*model.ProductResponse, error) {
	tx := service.Database.Preload("Images").Preload("Categories").WithContext(ctx)
	var products []*entity.Product
	for _, category := range categories {
		item := new(entity.Category)
		var results []*entity.Product
		if err := service.Database.Where("name = ?", category).First(item).Error; err != nil {
			return nil, err
		}
		if err := tx.Model(&item).Association("Products").Find(&results); err != nil {
			return nil, err
		}
		products = append(products, results...)
	}
	responses := make([]*model.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = converter.ToProductResponse(product)
	}
	return responses, nil
}

func (service *ProductService) GetProduct(ctx context.Context, id uint) (*model.ProductResponse, error) {
	product := new(entity.Product)
	tx := service.Database.Model(product).Preload("Categories").Preload("Images").WithContext(ctx)
	if err := service.ProductRepository.FindById(tx, id, product); err != nil {
		return nil, err
	}
	return converter.ToProductResponse(product), nil
}

func (service *ProductService) UpdateProduct(ctx context.Context, id uint, request *model.EditProductRequest) (*model.ProductResponse, error) {
	product := new(entity.Product)
	tx := service.Database.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := service.ProductRepository.FindById(tx, id, product); err != nil {
		return nil, err
	}
	product.Name = request.Name
	product.Name = request.Description
	product.Price = request.Price
	categories := make([]*entity.Category, len(request.Categories))
	for i, category := range request.Categories {
		item := new(entity.Category)
		if err := service.Database.Where("name = ?", category).First(item).Error; err != nil {
			service.Logger.Errorf("error on finding categories %+v", err)
			break
		}
		categories[i] = item
	}
	product.Categories = categories
	return converter.ToProductResponse(product), nil
}

func (service *ProductService) DeleteProduct(ctx context.Context, id uint) error {
	product := new(entity.Product)
	tx := service.Database.Model(product).Preload("Categories").Preload("Images").WithContext(ctx)
	if err := service.ProductRepository.FindById(tx, id, product); err != nil {
		return err
	}
	for _, image := range product.Images {
		service.Uploader.Upload.Destroy(context.Background(), uploader.DestroyParams{PublicID: image.PublicID})
	}
	if err := service.ProductRepository.Delete(tx, product); err != nil {
		return err
	}
	return nil
}
