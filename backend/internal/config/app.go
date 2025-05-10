package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rajaabluu/commerce/backend/internal/http/handler"
	"github.com/rajaabluu/commerce/backend/internal/http/route"
	"github.com/rajaabluu/commerce/backend/internal/repository"
	"github.com/rajaabluu/commerce/backend/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type App struct {
	Router    *echo.Echo
	Config    *viper.Viper
	Logger    *logrus.Logger
	Database  *gorm.DB
	Validator *validator.Validate
	Uploader  *cloudinary.Cloudinary
}

func (app *App) Init() {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(app.Config, app.Validator, app.Logger, app.Database, userRepository)
	userHandler := handler.NewUserHandler(app.Logger, userService)

	route := &route.RouteConfig{
		Route:       app.Router,
		UserHandler: userHandler,
	}

	route.Setup()
}
