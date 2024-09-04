package service

import (
	"context"
	"fmt"
	"strings"

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

func (service *ProductService) CreateProduct(ctx context.Context, request *model.CreateProductRequest) (*model.ProductResponse, error) {
	tx := service.Database.Model(&entity.Product{}).Preload("Categories").WithContext(ctx).Begin()
	defer tx.Rollback()
	product := &entity.Product{
		Name:        request.Name,
		Description: request.Description,
		Slug:        strings.ReplaceAll(request.Name, " ", "-"),
		Quantity:    request.Quantity,
	}
	categories := make([]*entity.Category, len(request.Categories))
	for i, id := range request.Categories {
		categories[i] = &entity.Category{ID: id}
	}
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

func (service *ProductService) GetProductsByCategory(ctx context.Context, IDs []uint) ([]*model.ProductResponse, error) {
	tx := service.Database.Preload("Images").WithContext(ctx)
	var products []*entity.Product
	categories := make([]*entity.Category, len(IDs))
	for i, ID := range IDs {
		categories[i] = &entity.Category{ID: ID}
	}
	if err := service.ProductRepository.Find(tx, &entity.Product{Categories: categories}, &products); err != nil {
		return nil, err
	}
	responses := make([]*model.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = converter.ToProductResponse(product)
	}
	return responses, nil
}

func (service *ProductService) GetProductBySlug(ctx context.Context, slug string) (*model.ProductResponse, error) {
	product := &entity.Product{Slug: slug}
	tx := service.Database.Model(product).Preload("Categories").Preload("Images").WithContext(ctx)
	if err := service.ProductRepository.FindOne(tx, product); err != nil {
		return nil, err
	}
	return converter.ToProductResponse(product), nil
}

func (service *ProductService) DeleteProductByID(ctx context.Context, ID uint) error {
	product := new(entity.Product)
	tx := service.Database.Model(product).Preload("Categories").Preload("Images").WithContext(ctx)
	if err := service.ProductRepository.FindByID(tx, ID, product); err != nil {
		return err
	}
	for _, image := range product.Images {
		service.Uploader.Upload.Destroy(context.Background(), uploader.DestroyParams{PublicID: image.PublicID})
	}
	if err := service.ProductRepository.DeleteByID(tx, ID); err != nil {
		return err
	}
	return nil
}
