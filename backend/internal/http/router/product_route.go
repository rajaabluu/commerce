package router

import (
	"github.com/labstack/echo/v4"
)

func (c *RouteConfig) SetupProductRoute(g *echo.Group) {
	product := g.Group("/products")
	product.GET("", c.ProductHandler.GetAll, c.Middleware.VerifyAuth)
	product.POST("", c.ProductHandler.CreateNewProduct, c.Middleware.VerifyAuth, c.Middleware.VerifyIsAdmin)
}
