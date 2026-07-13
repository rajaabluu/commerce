package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rajaabluu/commerce/backend/internal/http/handler"
	customMiddleware "github.com/rajaabluu/commerce/backend/internal/http/middleware"
	"github.com/sirupsen/logrus"
)

type RouterConfig struct {
	Route      *echo.Echo
	Logger     *logrus.Logger
	Middleware *customMiddleware.Middleware

	UserHandler         *handler.UserHandler
	ProductHandler      *handler.ProductHandler
	ProductImageHandler *handler.ProductImageHandler
	AddressHandler      *handler.AddressHandler
	OrderHandler        *handler.OrderHandler
	PaymentHandler      *handler.PaymentHandler
}

func (c *RouterConfig) Register() {
	c.Route.Use(middleware.Logger())
	c.Route.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello World!",
		})
	})

	api := c.Route.Group("/api")
	auth := api.Group("/auth")
	users := api.Group("/users", c.Middleware.VerifyAuth)
	products := api.Group("/products", c.Middleware.VerifyAuth)
	orders := api.Group("/orders", c.Middleware.VerifyAuth)
	payments := api.Group("/payments")

	c.registerAuthRouter(auth)
	c.registerUserRouter(users)
	c.registerProductRouter(products)
	c.registerOrderRouter(orders)
	c.registerPaymentRouter(payments)
}
