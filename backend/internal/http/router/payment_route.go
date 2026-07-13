package router

import "github.com/labstack/echo/v4"

func (c *RouterConfig) registerPaymentRouter(r *echo.Group) {
	r.POST("/notification", c.PaymentHandler.HandlePaymentNotification)
}
