package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rajaabluu/commerce/backend/internal/http/handler"
	customMiddleware "github.com/rajaabluu/commerce/backend/internal/http/middleware"
	"github.com/sirupsen/logrus"
)

type RouteConfig struct {
	Route          *echo.Echo
	UserHandler    *handler.UserHandler
	ProductHandler *handler.ProductHandler
	Logger         *logrus.Logger
	Middleware     *customMiddleware.Middleware
}

func (c *RouteConfig) Setup() {
	c.Route.Use(middleware.Logger())
	c.Route.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello World!",
		})
	})
	api := c.Route.Group("/api")
	c.SetupUserRoute(api)
	c.SetupProductRoute(api)
}
