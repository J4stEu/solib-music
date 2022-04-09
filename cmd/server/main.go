package main

import (
	"os"

	"github.com/J4stEu/solib/internal/app/config"
	"github.com/J4stEu/solib/internal/app/errors"
	"github.com/J4stEu/solib/internal/app/errors/server_errors"
	"github.com/J4stEu/solib/internal/app/server"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

func init() {
	logger = log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	if _, isDebug := os.LookupEnv("DEBUG"); isDebug {
		err := godotenv.Load("./configs/dev.env")
		if err != nil {
			logger.Fatal(errors.SetError(errors.ServerErrorLevel, server_errors.DevEnvFileNotFoundError, err))
		}
	}
	if !config.CheckENV() {
		logger.Fatal(errors.SetError(errors.ServerErrorLevel, server_errors.EnvSetError, nil))
	}
}

func main() {
	configuration, err := config.ReadConfiguration()
	if err != nil {
		logger.Fatal(err)
	}
	application := server.New(configuration, logger)
	if err = application.Start(); err != nil {
		logger.Fatal(err)
	}
}
