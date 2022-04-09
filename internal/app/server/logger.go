package server

import (
	"github.com/J4stEu/solib/internal/app/errors"
	"github.com/J4stEu/solib/internal/app/errors/server_errors"
	"github.com/sirupsen/logrus"
)

// ConfigureLogger - logger configuration
func (srv *Server) ConfigureLogger() error {
	srv.logger.Info("Configuring logger...")
	var logLevel logrus.Level
	switch srv.config.Server.LogLevel {
	case "debug":
		logLevel = logrus.DebugLevel
	case "info":
		logLevel = logrus.InfoLevel
	case "warn":
		logLevel = logrus.WarnLevel
	case "error":
		logLevel = logrus.ErrorLevel
	case "fatal":
		logLevel = logrus.FatalLevel
	default:
		return errors.SetError(errors.ServerErrorLevel, server_errors.LoggerLevelError, nil)
	}
	srv.logger.SetLevel(logLevel)
	return nil
}
