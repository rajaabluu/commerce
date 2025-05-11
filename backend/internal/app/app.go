package app

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rajaabluu/commerce/backend/internal/config"
	"github.com/rajaabluu/commerce/backend/internal/http/handler"
	"github.com/rajaabluu/commerce/backend/internal/http/middleware"
	"github.com/rajaabluu/commerce/backend/internal/http/router"
	"github.com/rajaabluu/commerce/backend/internal/repository"
	"github.com/rajaabluu/commerce/backend/internal/service"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type App struct {
	Router    *echo.Echo
	Config    *config.Config
	Logger    *logrus.Logger
	Database  *gorm.DB
	Validator *validator.Validate
	Uploader  *cloudinary.Cloudinary
}

func Bootstrap(c *App) {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(c.Config, c.Validator, c.Logger, c.Database, userRepository)
	userHandler := handler.NewUserHandler(c.Logger, userService)

	route := &router.RouteConfig{
		Route:       c.Router,
		UserHandler: userHandler,
		Logger:      c.Logger,
		Middleware:  middleware.NewMiddleware(c.Logger, c.Config),
	}

	route.Setup()
}
