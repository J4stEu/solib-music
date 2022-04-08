package server

import "github.com/sirupsen/logrus"

// ConfigureLogger - logger configuration
func (srv *Server) ConfigureLogger() {
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
		srv.logger.Fatal("Error setting logger format.")
	}
	srv.logger.SetLevel(logLevel)
}
