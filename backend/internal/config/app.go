package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/rajaabluu/ershop-api/internal/http/handler"
	"github.com/rajaabluu/ershop-api/internal/http/handler/middleware"
	"github.com/rajaabluu/ershop-api/internal/http/route"
	"github.com/rajaabluu/ershop-api/internal/repository"
	"github.com/rajaabluu/ershop-api/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Config struct {
	Database  *gorm.DB
	Router    *chi.Mux
	Validator *validator.Validate
	Uploader  *cloudinary.Cloudinary
	Config    *viper.Viper
	Logger    *logrus.Logger
}

func Init(config *Config) *chi.Mux {

	customerRepository := repository.NewUserRepository()
	productRepository := repository.NewProductRepository()
	orderRepository := repository.NewOrderRepository()
	paymentRepository := repository.NewPaymentRepository()

	customerService := service.NewUserService(config.Config, config.Database, config.Validator, config.Logger, customerRepository)
	productService := service.NewProductService(config.Database, config.Validator, config.Uploader, config.Logger, productRepository)
	paymentService := service.NewPaymentService(config.Logger, NewSnapClient(config.Config), config.Database, paymentRepository)
	orderService := service.NewOrderService(config.Config, config.Logger, config.Validator, config.Database, productService, paymentService, orderRepository)

	authHandler := handler.NewAuthHandler(config.Logger, customerService)
	productHandler := handler.NewProductHandler(config.Logger, productService)
	orderHandler := handler.NewOrderHandler(orderService)

	route := route.Config{
		Router:         config.Router,
		AuthHandler:    authHandler,
		ProductHandler: productHandler,
		OrderHandler:   orderHandler,
		Config:         config.Config,
		Middleware:     middleware.NewMiddleware(customerService),
	}

	route.Setup()

	return config.Router
}
