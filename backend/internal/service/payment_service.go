package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/rajaabluu/commerce/backend/internal/config"
	"github.com/rajaabluu/commerce/backend/internal/entity"
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
	tx := s.DB.WithContext(ctx).Begin()

	defer tx.Rollback()
	orderID, err := strconv.Atoi(req.OrderID)

	if err != nil {
		return err
	}

	payment, err := s.PaymentRepository.FindOne(tx, &entity.Payment{OrderID: uint(orderID)})

	if err != nil {
		return err
	}

	switch req.TransactionStatus {
	case "pending":
		payment.Status = entity.PaymentPending
	case "expire":
		payment.Status = entity.PaymentExpired
	case "settlement":
		payment.Status = entity.PaymentPaid
	case "deny":
		payment.Status = entity.PaymentFailed
	case "cancel":
		payment.Status = entity.PaymentCancelled
	case "refund", "chargeback":
		payment.Status = entity.PaymentRefunded
	}

	payment.Method = req.PaymentType
	payment.TransactionID = req.TransactionID

	if err := s.PaymentRepository.Update(tx, payment); err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
