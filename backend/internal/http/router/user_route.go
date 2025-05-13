package router

import "github.com/labstack/echo/v4"

func (c *RouteConfig) SetupUserRoute(g *echo.Group) {
	auth := g.Group("/auth")
	auth.POST("/register", c.UserHandler.Register)
	auth.POST("/login", c.UserHandler.Login)
	user := g.Group("/users")
	user.GET("/me", c.UserHandler.GetAuthenticatedUser, c.Middleware.VerifyAuth)
	user.PUT("/me", c.UserHandler.UpdateProfile, c.Middleware.VerifyAuth)
}
