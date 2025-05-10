package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rajaabluu/commerce/backend/internal/http/handler"
)

type RouteConfig struct {
	Route       *echo.Echo
	UserHandler *handler.UserHandler
}

func (c *RouteConfig) Setup() {
	c.Route.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello World!",
		})
	})
	c.SetupUserRoute()
}
