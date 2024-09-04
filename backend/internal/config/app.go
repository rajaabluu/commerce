package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/rajaabluu/ershop-api/internal/handler"
	"github.com/rajaabluu/ershop-api/internal/handler/middleware"
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

	customerRepository := repository.NewCustomerRepository()
	productRepository := repository.NewProductRepository()

	customerService := service.NewCustomerService(config.Config, config.Database, config.Validator, config.Logger, customerRepository)
	productService := service.NewProductService(config.Database, config.Validator, config.Uploader, config.Logger, productRepository)

	authHandler := handler.NewAuthHandler(config.Logger, customerService)
	productHandler := handler.NewProductHandler(config.Logger, productService)

	route := &RouteConfig{
		Router:         config.Router,
		AuthHandler:    authHandler,
		ProductHandler: productHandler,
		Config:         config.Config,
		Middleware:     middleware.NewMiddleware(customerService),
	}

	route.Setup()

	return config.Router
}
