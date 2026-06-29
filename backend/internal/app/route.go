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
	productService := service.NewProductService(
		app.Config,
		app.Logger,
		app.Database,
		app.Validator,
		productRepository,
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
