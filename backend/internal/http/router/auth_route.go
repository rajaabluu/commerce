package router

import "github.com/labstack/echo/v4"

func (c *RouterConfig) registerAuthRouter(r *echo.Group) {
	r.POST("/register", c.UserHandler.Register)
	r.POST("/login", c.UserHandler.Login)
}
