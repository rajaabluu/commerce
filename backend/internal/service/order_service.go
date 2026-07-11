package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/rajaabluu/commerce/backend/internal/config"
	"github.com/rajaabluu/commerce/backend/internal/entity"
	"github.com/rajaabluu/commerce/backend/internal/helper/mapper"
	"github.com/rajaabluu/commerce/backend/internal/model"
	"github.com/rajaabluu/commerce/backend/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderService struct {
	Config *config.Config
	Logger *logrus.Logger

	DB         *gorm.DB
	Validator  *validator.Validate
	PaymentLib snap.Client

	ProductRepository *repository.ProductRepository
	PaymentRepository *repository.PaymentRepository
	UserRepository    *repository.UserRepository
	AddressRepository *repository.AddressRepository
	OrderRepository   *repository.OrderRepository
}

func NewOrderService(
	config *config.Config,
	logger *logrus.Logger,

	DB *gorm.DB,
	validator *validator.Validate,
	paymentLib snap.Client,

	productRepository *repository.ProductRepository,
	paymentRepository *repository.PaymentRepository,
	userRepository *repository.UserRepository,
	addressRepository *repository.AddressRepository,
	orderRepository *repository.OrderRepository,
) *OrderService {
	return &OrderService{
		Config: config,
		Logger: logger,

		DB:         DB,
		Validator:  validator,
		PaymentLib: paymentLib,

		ProductRepository: productRepository,
		PaymentRepository: paymentRepository,
		UserRepository:    userRepository,
		AddressRepository: addressRepository,
		OrderRepository:   orderRepository,
	}
}

func (s *OrderService) Create(
	ctx context.Context,
	userID uint,
	req *model.CreateOrderRequest,
) (*model.OrderResponse, error) {

	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(req); err != nil {
		return nil, err
	}

	user, err := s.UserRepository.FindById(tx, userID)

	if err != nil {
		return nil, err
	}

	address, err := s.AddressRepository.FindById(tx, req.ShippingAddressID)
	if err != nil {
		return nil, err
	}

	productIDs := make([]uint, 0, len(req.OrderItems))

	for _, item := range req.OrderItems {
		productIDs = append(productIDs, item.ProductID)
	}

	products, err := s.ProductRepository.FindByIDs(tx, productIDs)
	if err != nil {
		return nil, err
	}

	productMap := make(map[uint]*entity.Product, len(products))

	for _, product := range products {
		productMap[product.ID] = product
	}

	var (
		totalPrice int64
		orderItems = make([]entity.OrderItem, 0, len(req.OrderItems))
	)

	for _, item := range req.OrderItems {

		product, ok := productMap[item.ProductID]
		if !ok {
			return nil, fmt.Errorf("product %d not found", item.ProductID)
		}

		if product.Stock < uint(item.Quantity) {
			return nil, fmt.Errorf("%s stock is insufficient", product.Name)
		}

		subtotal := int64(product.Price) * int64(item.Quantity)

		totalPrice += subtotal

		orderItems = append(orderItems, entity.OrderItem{
			ProductID:   product.ID,
			ProductName: product.Name,
			Price:       int64(product.Price),
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})

		product.Stock -= uint(item.Quantity)
	}

	invoiceID := fmt.Sprintf(
		"INV-%s-%06d",
		time.Now().Format("20060102"),
		rand.Intn(1000000),
	)

	order := &entity.Order{
		UserID:        userID,
		InvoiceID:     invoiceID,
		Status:        entity.OrderPending,
		RecipientName: address.RecipientName,
		Phone:         address.Phone,
		Province:      address.Province,
		City:          address.City,
		District:      address.District,
		PostalCode:    address.PostalCode,
		StreetAddress: address.StreetAddress,

		TotalPrice: totalPrice,
		OrderItems: orderItems,
	}

	if err := s.OrderRepository.Create(tx, order); err != nil {
		return nil, err
	}

	for _, product := range products {
		if err := s.ProductRepository.Update(tx, product); err != nil {
			return nil, err
		}
	}

	itemDetails := make([]midtrans.ItemDetails, 0, len(order.OrderItems))

	for _, item := range order.OrderItems {
		itemDetails = append(itemDetails, mapper.ToMidtransItem(item))
	}

	snapReq := mapper.ToSnapRequest(order, user, address)
	snapReq.Items = &itemDetails

	snapRes, err := s.PaymentLib.CreateTransaction(snapReq)

	payment := &entity.Payment{
		OrderID:     order.ID,
		Token:       snapRes.Token,
		RedirectURL: snapRes.RedirectURL,
	}

	if err := s.PaymentRepository.Create(tx, payment); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	items := make([]*model.OrderItemResponse, 0, len(order.OrderItems))

	for _, item := range order.OrderItems {
		items = append(items, &model.OrderItemResponse{
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Price:       item.Price,
			Quantity:    uint(item.Quantity),
			Subtotal:    item.Subtotal,
		})
	}

	paymentRes := &model.Payment{
		OrderID:     order.ID,
		Token:       snapRes.Token,
		RedirectURL: snapRes.RedirectURL,
	}

	return &model.OrderResponse{
		ID:            order.ID,
		Status:        string(order.Status),
		TotalPrice:    order.TotalPrice,
		RecipientName: order.RecipientName,
		Phone:         order.Phone,
		Province:      order.Province,
		City:          order.City,
		District:      order.District,
		PostalCode:    order.PostalCode,
		StreetAddress: order.StreetAddress,
		OrderItems:    items,
		Payment:       *paymentRes,
	}, nil
}
