package config

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rajaabluu/ershop-api/internal/handler"
	"github.com/rajaabluu/ershop-api/internal/handler/middleware"
	"github.com/spf13/viper"
)

type RouteConfig struct {
	Router         *chi.Mux
	Config         *viper.Viper
	AuthHandler    *handler.AuthHandler
	ProductHandler *handler.ProductHandler
	Middleware     *middleware.Middleware
}

func (config *RouteConfig) Setup() {
	config.Router.Use(chiMiddleware.Logger)
	config.Router.Route("/api", func(r chi.Router) {
		config.SetupAuthHandler(r)
		config.SetupProductHandler(r)
	})
}

func (config *RouteConfig) SetupAuthHandler(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", config.AuthHandler.Register)
		r.Post("/login", config.AuthHandler.Login)
	})
}

func (config *RouteConfig) SetupProductHandler(r chi.Router) {
	r.Route("/product", func(r chi.Router) {
		r.Use(config.Middleware.VerifyAuth)
		r.Get("/", config.ProductHandler.GetAllProducts)
		r.Post("/", config.ProductHandler.CreateNewProduct)
		r.Get("/{id}", config.ProductHandler.GetProductDetail)
		r.Delete("/{id}", config.ProductHandler.DeleteProduct)
	})
}
