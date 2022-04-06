package app

import "github.com/sirupsen/logrus"

// App - application structure
type App struct {
	config *Config
	logger *logrus.Logger
}

// New - new application instance
func New(config *Config, logger *logrus.Logger) *App {
	return &App{
		config: config,
		logger: logger,
	}
}

// Start - start application instance
func (app *App) Start() error {
	app.logger.Info("Starting application...")
	return nil
}

// ConfigureLogger - logger configuration
func (app *App) ConfigureLogger() {
	var logLevel logrus.Level
	switch app.config.LogLevel {
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
		app.logger.Fatal("Error setting logger format.")
	}
	app.logger.SetLevel(logLevel)
}
