package service

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rajaabluu/ershop-api/internal/entity"
	"github.com/rajaabluu/ershop-api/internal/model"
	"github.com/rajaabluu/ershop-api/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type OrderService struct {
	Config          *viper.Viper
	Logger          *logrus.Logger
	Validator       *validator.Validate
	Database        *gorm.DB
	ProductService  *ProductService
	PaymentService  *PaymentService
	OrderRepository *repository.OrderRepository
}

func NewOrderService(viper *viper.Viper, logger *logrus.Logger, validator *validator.Validate, database *gorm.DB, productService *ProductService, paymentService *PaymentService, repository *repository.OrderRepository) *OrderService {
	return &OrderService{
		Config:          viper,
		Logger:          logger,
		Database:        database,
		ProductService:  productService,
		PaymentService:  paymentService,
		OrderRepository: repository,
		Validator:       validator,
	}
}

func (service *OrderService) CreateOrder(ctx context.Context, request *model.CreateOrderRequest) (*model.CreateOrderResponse, error) {
	tx := service.Database.WithContext(ctx).Begin()
	auth := ctx.Value(model.AuthContextKey).(*model.Auth)
	defer tx.Rollback()
	if err := service.Validator.Struct(request); err != nil {
		return nil, err
	}
	order := &entity.Order{UserID: auth.ID, ID: uuid.NewString()}
	if err := service.OrderRepository.Create(tx, order); err != nil {
		return nil, fmt.Errorf("error on creating order %+v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("error on commit transaction: %+v", err)
	}
	products := make([]*model.ProductResponse, len(request.Products))
	items := make([]*model.CreatePaymentProductRequest, len(request.Products))
	for i, product := range request.Products {
		result, err := service.ProductService.GetProduct(ctx, product.ID)
		if err != nil {
			return nil, err
		}
		products[i] = result
		items[i] = &model.CreatePaymentProductRequest{ID: result.ID, Price: result.Price, Name: result.Name, Quantity: result.Quantity}
	}
	response, err := service.PaymentService.CreatePayment(ctx, &model.CreatePaymentRequest{OrderID: order.ID, Products: items})
	if err != nil {
		return nil, err
	}
	return &model.CreateOrderResponse{RedirectUrl: response.RedirectUrl}, nil
}

// func (service *OrderService) CreateOrderDetails(ctx context.Context, request *model.CreateOrderDetailRequest) error {
// 	tx := service.Database.WithContext(ctx).Begin()
// 	defer tx.Rollback()
// 	item := &entity.OrderDetail{
// 		OrderID:   request.OrderID,
// 		ProductID: request.ProductID,
// 		Quantity:  request.Quantity,
// 		Total:     request.Total}
// 	if err := tx.Model(new(entity.OrderDetail)).Save(item).Error; err != nil {
// 		return fmt.Errorf("error on creating product detail: %+v", err)
// 	}
// 	if err := tx.Commit().Error; err != nil {
// 		return fmt.Errorf("error on commit transaction: %+v", err)
// 	}
// 	return nil
// }
