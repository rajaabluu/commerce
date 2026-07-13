package service

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/rajaabluu/commerce/backend/internal/config"
	"github.com/rajaabluu/commerce/backend/internal/model"
	"github.com/rajaabluu/commerce/backend/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PaymentService struct {
	Config *config.Config
	Logger *logrus.Logger

	DB        *gorm.DB
	Validator *validator.Validate

	PaymentRepository *repository.PaymentRepository
}

func NewPaymentService(

	config *config.Config,
	logger *logrus.Logger,

	db *gorm.DB,
	validator *validator.Validate,

	paymentRepository *repository.PaymentRepository,
) *PaymentService {
	return &PaymentService{
		Config:            config,
		Logger:            logger,
		DB:                db,
		Validator:         validator,
		PaymentRepository: paymentRepository,
	}
}

func (s *PaymentService) Create(ctx context.Context, userID uint, req any) error {
	return errors.New("not implemented yet")
}

func (s *PaymentService) Notification(ctx context.Context, req *model.MidtransNotificationRequest) error {
	return errors.New("not implemented yet")
}
