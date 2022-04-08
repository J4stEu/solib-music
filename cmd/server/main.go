package main

import (
	"github.com/J4stEu/solib/internal/app/config"
	"github.com/J4stEu/solib/internal/app/server"
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
	if !config.CheckENV() {
		logger.Fatal("Error setting environment.")
	}
}

func main() {
	application := server.New(config.ReadConfiguration(logger), logger)
	if err := application.Start(); err != nil {
		log.Fatal(err)
	}
}
