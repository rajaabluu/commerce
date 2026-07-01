package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/rajaabluu/commerce/backend/internal/config"
	"github.com/rajaabluu/commerce/backend/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderService struct {
	Config *config.Config
	Logger *logrus.Logger

	DB        *gorm.DB
	Validator *validator.Validate

	OrderRepository *repository.OrderRepository
}

func NewOrderService(
	config *config.Config,
	logger *logrus.Logger,

	DB *gorm.DB,
	validator *validator.Validate,

	orderRepository *repository.OrderRepository,
) *OrderService {
	return &OrderService{
		Config:          config,
		Logger:          logger,
		DB:              DB,
		Validator:       validator,
		OrderRepository: orderRepository,
	}
}
