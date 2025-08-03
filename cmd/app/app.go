package app

import (
	"image/internal/config"
	"image/internal/database"
	"image/internal/server"
	"log"
)

func Setup() {
	database.Connect()
	config.AutoMigrate()

	server.Setup()
	app := server.New()

	if err := app.Listen(":3015"); err != nil {
		log.Fatalf("unable to listen on port 3015 : %v", err)
	}
}
