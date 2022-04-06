package main

import (
	"github.com/J4stEu/solib/internal/app"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	if _, isDebug := os.LookupEnv("DEBUG"); isDebug {
		err := godotenv.Load("./configs/dev.env")
		if err != nil {
			logger.Fatal("Error loading dev.env file.")
		}
	}
	if !app.CheckENV() {
		logger.Fatal("Error setting environment.")
	}
}

func main() {
	config := app.ReadConfiguration(logger)
	application := app.New(config, logger)
	application.ConfigureLogger()
	if err := application.Start(); err != nil {
		log.Fatal(err)
	}
}
