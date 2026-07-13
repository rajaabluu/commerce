package router

import "github.com/labstack/echo/v4"

func (c *RouterConfig) registerOrderRouter(r *echo.Group) {
	r.POST("", c.OrderHandler.CreateNewOrder)
}
