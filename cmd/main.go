package main

import (
	"image/cmd/app"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("unable to load env")
	}

	app.Setup()
}
