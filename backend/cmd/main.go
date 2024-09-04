package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/rajaabluu/ershop-api/internal/config"
)

var app *config.Config
var PORT int

func init() {
	viper := config.NewViper()
	PORT = viper.GetInt("app.port")
	app = &config.Config{
		Database:  config.NewDatabase(viper),
		Router:    chi.NewRouter(),
		Validator: validator.New(),
		Logger:    config.NewLogger(),
		Config:    viper,
		Uploader:  config.NewUploader(viper),
	}
}

func main() {
	app := config.Init(app)
	log.Printf("starting server on http://localhost:%d", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), app)
	if err != nil {
		log.Fatalf("error on starting server : %v", err)
	}
}
