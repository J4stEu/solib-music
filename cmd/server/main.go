package main

import (
	"github.com/J4stEu/solib/internal/app/config"
	"github.com/J4stEu/solib/internal/app/errors"
	"github.com/J4stEu/solib/internal/app/errors/server_errors"
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
			logger.WithFields(log.Fields{
				"error": err,
			}).Fatal(errors.SetError(errors.ServerErrorLevel, server_errors.DevEnvFileNotFound))
		}
	}
	if !config.CheckENV() {
		logger.WithFields(log.Fields{
			"error": "Error setting environment.",
		}).Fatal(errors.SetError(errors.ServerErrorLevel, server_errors.EnvSetError))
	}
}

func main() {
	application := server.New(config.ReadConfiguration(logger), logger)
	if err := application.Start(); err != nil {
		logger.Fatal(err)
	}
}
