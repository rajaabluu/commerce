package main

import (
	"log"

	"github.com/rajaabluu/commerce/backend/internal/app"
)

func main() {
	app := app.NewApp()
	log.Fatal(app.Start(app.Config.App.Port))
}
