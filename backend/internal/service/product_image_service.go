package service

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/rajaabluu/commerce/backend/internal/config"
	"github.com/rajaabluu/commerce/backend/internal/entity"
	"github.com/rajaabluu/commerce/backend/internal/exception"
	"github.com/rajaabluu/commerce/backend/internal/model"
	"github.com/rajaabluu/commerce/backend/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductImageService struct {
	Config *config.Config
	Logger *logrus.Logger

	DB        *gorm.DB
	Validator *validator.Validate
	Uploader  *cloudinary.Cloudinary

	ProductImageRepository *repository.ProductImageRepository
}

func NewProductImageService(
	config *config.Config,
	logger *logrus.Logger,
	db *gorm.DB,
	validator *validator.Validate,
	uploader *cloudinary.Cloudinary,
	productImageRepository *repository.ProductImageRepository) *ProductImageService {

	return &ProductImageService{
		Config: config,
		Logger: logger,

		DB:        db,
		Validator: validator,
		Uploader:  uploader,

		ProductImageRepository: productImageRepository,
	}
}

func (s *ProductImageService) UploadProductImage(ctx context.Context, files []*multipart.FileHeader, productID uint) ([]*model.ProductImageResponse, error) {
	var productImages []*entity.ProductImage
	var publicIDs []string

	if len(files) == 0 {
		return nil, exception.ErrNoProductImages
	}

	if len(files) > 5 {
		return nil, exception.ErrTooManyProductImages
	}

	for _, fileHeader := range files {
		file, err := fileHeader.Open()

		if err != nil {
			return nil, err
		}

		defer file.Close()

		cldRes, err := s.Uploader.Upload.Upload(ctx, file, uploader.UploadParams{
			Folder: "/commerce/products",
		})

		if err != nil {
			s.Logger.Warnf("failed to upload product image: %+v", err)
			return nil, err
		}

		publicIDs = append(publicIDs, cldRes.PublicID)

		productImage := &entity.ProductImage{
			PublicID:  cldRes.PublicID,
			ProductID: productID,
			Source:    cldRes.SecureURL,
		}

		productImages = append(productImages, productImage)
	}

	tx := s.DB.Begin().WithContext(ctx)

	defer tx.Rollback()
	if err := s.ProductImageRepository.BulkCreate(tx, productImages); err != nil {
		s.Logger.Warnf("failed to create product image: %+v", err)
		for _, id := range publicIDs {
			s.Uploader.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: id})
		}
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		s.Logger.Warnf("failed to create commit transaction: %+v", err)
		return nil, err
	}

	var res []*model.ProductImageResponse

	for _, productImage := range productImages {
		res = append(res, &model.ProductImageResponse{
			ID:        productImage.ID,
			ProductID: productImage.ProductID,
			PublicID:  productImage.PublicID,
			Source:    productImage.Source,
		})
	}

	return res, nil

}
