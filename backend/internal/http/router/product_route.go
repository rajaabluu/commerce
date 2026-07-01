package router

import (
	"github.com/labstack/echo/v4"
)

func (c *RouterConfig) registerProductRouter(r *echo.Group) {
	r.GET("", c.ProductHandler.GetAll)
	r.POST("", c.ProductHandler.CreateNewProduct, c.Middleware.VerifyIsAdmin)
	r.GET("/:id", c.ProductHandler.GetProductById)
	r.PUT("/:id", c.ProductHandler.UpdateProduct)
	r.POST("/:id/images", c.ProductImageHandler.UploadProductImage, c.Middleware.VerifyIsAdmin)
	r.DELETE("/:id", c.ProductHandler.DeleteProduct)
}
