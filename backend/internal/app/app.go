package app

import (
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/rajaabluu/commerce/backend/internal/config"
	"github.com/rajaabluu/commerce/backend/internal/http/middleware"
	"github.com/rajaabluu/commerce/backend/internal/http/router"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type App struct {
	Router     *echo.Echo
	Config     *config.Config
	Logger     *logrus.Logger
	Database   *gorm.DB
	Validator  *validator.Validate
	Uploader   *cloudinary.Cloudinary
	PaymentLib snap.Client
}

func NewApp() *App {
	r := echo.New()
	cfg := config.NewConfig()
	logger := config.NewLogger()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	app := &App{
		Router:     r,
		Config:     cfg,
		Logger:     logger,
		Database:   config.NewDatabase(cfg),
		Uploader:   config.NewUploader(cfg),
		Validator:  validator.New(),
		PaymentLib: config.NewConfig().NewPaymentLib(),
	}

	userHandler := registerUserHandler(app)
	productHandler := registerProductHandler(app)
	productImageHandler := registerProductImageHandler(app)

	routeCfg := &router.RouterConfig{
		Route:               app.Router,
		UserHandler:         userHandler,
		ProductHandler:      productHandler,
		ProductImageHandler: productImageHandler,
		Logger:              app.Logger,
		Middleware:          middleware.NewMiddleware(app.Logger, app.Config),
	}

	routeCfg.Register()

	return app
}

func (a *App) Start(port int) error {
	a.Logger.Printf("server started on http://localhost:%d", a.Config.App.Port)
	return a.Router.Start(fmt.Sprintf(":%d", port))
}
