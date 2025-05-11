package router

import "github.com/labstack/echo/v4"

func (c *RouteConfig) SetupUserRoute(g *echo.Group) {
	r := g.Group("/auth")
	r.POST("/register", c.UserHandler.Register)
	r.POST("/login", c.UserHandler.Login)
	r.GET("/me", c.UserHandler.GetAuthenticatedUser, c.Middleware.VerifyAuth)
}
