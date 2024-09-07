package service

import (
	"github.com/rajaabluu/ershop-api/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type OrderService struct {
	Config          *viper.Viper
	Logger          *logrus.Logger
	Database        *gorm.DB
	OrderRepository repository.OrderRepository
}

func NewOrderService(viper *viper.Viper, logger *logrus.Logger, database *gorm.DB, repository repository.OrderRepository) *OrderService {
	return &OrderService{
		Config:          viper,
		Logger:          logger,
		Database:        database,
		OrderRepository: repository,
	}
}
