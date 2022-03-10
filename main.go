package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/veezyjay/banka/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	app.Start()
}
