package route

func (c *RouteConfig) SetupUserRoute() {
	c.Route.POST("/register", c.UserHandler.Register)
}
