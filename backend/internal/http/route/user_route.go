package route

func (c *RouteConfig) SetupUserRoute() {
	c.Route.POST("/auth/register", c.UserHandler.Register)
}
