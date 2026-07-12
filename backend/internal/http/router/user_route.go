package router

import "github.com/labstack/echo/v4"

func (c *RouterConfig) registerUserRouter(r *echo.Group) {
	r.GET("/me", c.UserHandler.GetAuthenticatedUser)
	r.PUT("/me", c.UserHandler.UpdateProfile)
	r.GET("/addresses", c.AddressHandler.GetAllAddresses)
	r.POST("/addresses", c.AddressHandler.CreateNewAddress)
}
