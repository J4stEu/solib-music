package app

import (
	"github.com/J4stEu/solib/internal/pkg"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

// Config - application configuration
type Config struct {
	// Server configuration
	ServerAddr string
	ServerPort uint

	// Database Configuration
	PostgresIP   string
	PostgresPort uint
	PostgresDB   string
	PostgresUser string
	PostgresPass string

	// Logging
	LogLevel string
}

func CheckENV() bool {
	_, err := os.LookupEnv("SERVER_ADDR")
	if !err {
		return false
	}
	_, err = os.LookupEnv("SERVER_PORT")
	if !err {
		return false
	}
	_, err = os.LookupEnv("PG_IP")
	if !err {
		return false
	}
	_, err = os.LookupEnv("PG_PORT")
	if !err {
		return false
	}
	_, err = os.LookupEnv("PG_DATABASE")
	if !err {
		return false
	}
	_, err = os.LookupEnv("PG_USER")
	if !err {
		return false
	}
	_, err = os.LookupEnv("PG_PASSWORD")
	if !err {
		return false
	}
	_, err = os.LookupEnv("PG_PASSWORD")
	if !err {
		return false
	}
	_, err = os.LookupEnv("LOG_LEVEL")
	if !err {
		return false
	}
	return true
}

func ReadConfiguration(logger *logrus.Logger) *Config {
	config := &Config{}
	// Server configuration
	// ServerAddr
	serverAddr, err := os.LookupEnv("SERVER_ADDR")
	if !err {
		logger.Fatal(err)
	}
	validServerAddr := pkg.IsValidIP(serverAddr)
	if !validServerAddr {
		logger.Fatal("Invalid server IP address.")
	}
	// ServerPort
	config.ServerAddr = serverAddr
	serverPort, err := os.LookupEnv("SERVER_PORT")
	if !err {
		logger.Fatal(err)
	}
	serverPortUINT, convertErr := strconv.Atoi(serverPort)
	if convertErr != nil {
		logger.Fatal(convertErr)
	}
	config.ServerPort = uint(serverPortUINT)

	// Database Configuration
	// PostgresIP
	postgresIP, err := os.LookupEnv("PG_IP")
	if !err {
		logger.Fatal(err)
	}
	validPostgresIP := pkg.IsValidIP(postgresIP)
	if !validPostgresIP {
		logger.Fatal("Invalid postgres IP address.")
	}
	config.PostgresIP = postgresIP
	// PostgresPort
	postgresPort, err := os.LookupEnv("PG_PORT")
	if !err {
		logger.Fatal(err)
	}
	postgresPortUINT, convertErr := strconv.Atoi(postgresPort)
	if convertErr != nil {
		logger.Fatal(convertErr)
	}
	config.PostgresPort = uint(postgresPortUINT)
	// PostgresDB
	postgresDB, err := os.LookupEnv("PG_DATABASE")
	if !err {
		logger.Fatal(err)
	}
	config.PostgresDB = postgresDB
	// PostgresUser
	postgresUser, err := os.LookupEnv("PG_USER")
	if !err {
		logger.Fatal(err)
	}
	config.PostgresUser = postgresUser
	// PostgresPass
	postgresPass, err := os.LookupEnv("PG_PASSWORD")
	if !err {
		logger.Fatal(err)
	}
	config.PostgresPass = postgresPass
	// LogLevel
	logLevel, err := os.LookupEnv("LOG_LEVEL")
	if !err {
		logger.Fatal(err)
	}
	config.LogLevel = logLevel
	return config
}
