package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/rajaabluu/ershop-api/internal/entity"
	"github.com/rajaabluu/ershop-api/internal/model"
	"github.com/rajaabluu/ershop-api/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PaymentService struct {
	SnapClient        *snap.Client
	Logger            *logrus.Logger
	Database          *gorm.DB
	PaymentRepository *repository.PaymentRepository
}

func NewPaymentService(logger *logrus.Logger, snapClient *snap.Client, database *gorm.DB, repository *repository.PaymentRepository) *PaymentService {
	return &PaymentService{
		Logger:            logger,
		SnapClient:        snapClient,
		Database:          database,
		PaymentRepository: repository,
	}
}

func (service *PaymentService) CreatePayment(ctx context.Context, request *model.CreatePaymentRequest) (*model.CreatePaymentResponse, error) {
	payment := new(entity.Payment)
	tx := service.Database.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := service.PaymentRepository.Create(tx, payment); err != nil {
		return nil, err
	}
	if err := tx.Model(&entity.Order{}).Where("ID = ?", request.OrderID).Update("payment_id", payment.ID).Error; err != nil {
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("error on commit transaction: %+v", err)
	}
	total := 0
	items := make([]midtrans.ItemDetails, len(request.Products))
	for i, product := range request.Products {
		item := midtrans.ItemDetails{Name: product.Name, Qty: int32(product.Quantity), ID: strconv.Itoa(int(product.ID)), Price: int64(product.Price)}
		items[i] = item
		total += int(product.Price) * int(product.Quantity)
	}
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  request.OrderID,
			GrossAmt: int64(total),
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items:           &items,
	}
	res, err := service.SnapClient.CreateTransaction(snapReq)
	if err != nil {
		return nil, err
	}
	return &model.CreatePaymentResponse{RedirectUrl: res.RedirectURL}, nil
}
