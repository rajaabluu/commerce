package router

import (
	"github.com/labstack/echo/v4"
)

func (c *RouteConfig) SetupProductRoute(g *echo.Group) {
	g.POST("/products", c.ProductHandler.CreateNewProduct, c.Middleware.VerifyAuth, c.Middleware.VerifyIsAdmin)
}
