package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rajaabluu/commerce/backend/internal/app"
	"github.com/rajaabluu/commerce/backend/internal/config"
)

func main() {
	c := echo.New()
	cfg := config.NewConfig()
	PORT := cfg.App.Port
	app.Bootstrap(&app.App{
		Router:    c,
		Config:    cfg,
		Logger:    config.NewLogger(),
		Database:  config.NewDatabase(cfg),
		Uploader:  config.NewUploader(cfg),
		Validator: validator.New(),
	})
	log.Printf("server started on http://localhost:%d", PORT)
	c.Logger.Fatal(c.Start(fmt.Sprintf(":%d", PORT)))

}
