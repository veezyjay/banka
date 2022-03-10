package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/veezyjay/banka/app"
	"github.com/veezyjay/banka/logger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	logger.Info("Starting the application...")
	app.Start()
}
