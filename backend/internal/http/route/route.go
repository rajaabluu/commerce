package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rajaabluu/ershop-api/internal/http/handler"
	"github.com/rajaabluu/ershop-api/internal/http/handler/middleware"
	"github.com/spf13/viper"
)

type Config struct {
	Router         *chi.Mux
	Config         *viper.Viper
	AuthHandler    *handler.AuthHandler
	ProductHandler *handler.ProductHandler
	OrderHandler   *handler.OrderHandler
	Middleware     *middleware.Middleware
}

func (config *Config) Setup() {
	config.Router.Use(chiMiddleware.Logger)
	config.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})
	config.Router.Route("/api", func(r chi.Router) {
		config.SetupAuthHandler(r)
		config.SetupProductHandler(r)
		config.SetupOrderHandler(r)
	})
}

func (config *Config) SetupAuthHandler(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/google", config.AuthHandler.UseGoogleAuth)
		r.Post("/register", config.AuthHandler.Register)
		r.Post("/login", config.AuthHandler.Login)
		r.Group(func(r chi.Router) {
			r.Use(config.Middleware.VerifyAuth)
			r.Get("/me", config.AuthHandler.GetMyProfile)
		})
	})
}

func (config *Config) SetupProductHandler(r chi.Router) {
	r.Route("/product", func(r chi.Router) {
		r.Use(config.Middleware.VerifyAuth)
		r.Get("/", config.ProductHandler.GetAllProducts)
		r.Post("/", config.ProductHandler.CreateNewProduct)
		r.Get("/categories", config.ProductHandler.GetProductCategories)
		r.Group(func(r chi.Router) {
			r.Use(config.Middleware.VerifyIsAdmin)
			r.Get("/{id}", config.ProductHandler.GetProductDetail)
			r.Delete("/{id}", config.ProductHandler.DeleteProduct)
		})
	})
}

func (config *Config) SetupOrderHandler(r chi.Router) {
	r.Route("/order", func(r chi.Router) {
		r.Use(config.Middleware.VerifyAuth)
		r.Post("/", config.OrderHandler.NewOrder)
	})
}
