package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rajaabluu/commerce/backend/internal/config"
)

func main() {
	c := echo.New()
	viper := config.NewViper()
	PORT := viper.GetInt("app.port")
	app := &config.App{
		Router:    c,
		Config:    viper,
		Logger:    config.NewLogger(),
		Database:  config.NewDatabase(viper),
		Uploader:  config.NewUploader(viper),
		Validator: validator.New(),
	}
	app.Init()
	log.Printf("server started on http://localhost:%d", PORT)
	c.Logger.Fatal(c.Start(fmt.Sprintf(":%d", PORT)))

}
