package app

import (
	"github.com/rajaabluu/commerce/backend/internal/http/handler"
	"github.com/rajaabluu/commerce/backend/internal/repository"
	"github.com/rajaabluu/commerce/backend/internal/service"
)

func registerUserHandler(app *App) *handler.UserHandler {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(
		app.Config,
		app.Logger,
		app.Database,
		app.Validator,
		userRepository,
	)
	userHandler := handler.NewUserHandler(app.Logger, userService)

	return userHandler
}

func registerProductHandler(app *App) *handler.ProductHandler {
	productRepository := repository.NewProductRepository()
	productImagerepository := repository.NewProductImageRepository()
	productService := service.NewProductService(
		app.Config,
		app.Logger,
		app.Database,
		app.Validator,
		app.Uploader,
		productRepository,
		productImagerepository,
	)
	productHandler := handler.NewProductHandler(app.Logger, productService)

	return productHandler
}

func registerProductImageHandler(app *App) *handler.ProductImageHandler {
	productImageRepository := repository.NewProductImageRepository()
	productImageService := service.NewProductImageService(
		app.Config,
		app.Logger,
		app.Database,
		app.Validator,
		app.Uploader,
		productImageRepository,
	)
	productImageHandler := handler.NewProductImageHandler(app.Logger, productImageService)

	return productImageHandler
}

func registerAddressHandler(app *App) *handler.AddressHandler {
	addressRepository := repository.NewAddressRepository()
	addressService := service.NewAddressService(app.Config, app.Logger, app.Database, app.Validator, addressRepository)
	addressHandler := handler.NewAddressHandler(app.Logger, addressService)

	return addressHandler
}

func registerOrderHandler(app *App) *handler.OrderHandler {
	productRepository := repository.NewProductRepository()
	paymentRepository := repository.NewPaymentRepository()
	orderRepository := repository.NewOrderRepository()
	userRepository := repository.NewUserRepository()
	addressRepository := repository.NewAddressRepository()
	orderService := service.NewOrderService(
		app.Config,
		app.Logger,
		app.Database,
		app.Validator,
		app.PaymentLib,

		productRepository,
		paymentRepository,
		userRepository,
		addressRepository,
		orderRepository,
	)

	orderHandler := handler.NewOrderHandler(app.Logger, orderService)

	return orderHandler
}

func registerPaymentHandler(app *App) *handler.PaymentHandler {
	paymentRepository := repository.NewPaymentRepository()
	paymentService := service.NewPaymentService(
		app.Config,
		app.Logger,
		app.Database,
		app.Validator,
		paymentRepository,
	)

	paymentHandler := handler.NewPaymentHandler(app.Logger, paymentService)

	return paymentHandler
}
