package main

import (
	"github.com/veezyjay/banka/app"
	"github.com/veezyjay/banka/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
