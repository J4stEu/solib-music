package config

import (
	"github.com/J4stEu/solib/internal/pkg"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

// Server - server configuration
type Server struct {
	ServerAddr string
	ServerPort uint
	// Logging
	LogLevel string
}

// DataBase - database configuration
type DataBase struct {
	PostgresIP   string
	PostgresPort uint
	PostgresDB   string
	PostgresUser string
	PostgresPass string
}

// Config - application configuration
type Config struct {
	Server   *Server
	DataBase *DataBase
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
	config := &Config{&Server{}, &DataBase{}}
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
	config.Server.ServerAddr = serverAddr
	serverPort, err := os.LookupEnv("SERVER_PORT")
	if !err {
		logger.Fatal(err)
	}
	serverPortUINT, convertErr := strconv.Atoi(serverPort)
	if convertErr != nil {
		logger.Fatal(convertErr)
	}
	config.Server.ServerPort = uint(serverPortUINT)
	// LogLevel
	logLevel, err := os.LookupEnv("LOG_LEVEL")
	if !err {
		logger.Fatal(err)
	}
	config.Server.LogLevel = logLevel

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
	config.DataBase.PostgresIP = postgresIP
	// PostgresPort
	postgresPort, err := os.LookupEnv("PG_PORT")
	if !err {
		logger.Fatal(err)
	}
	postgresPortUINT, convertErr := strconv.Atoi(postgresPort)
	if convertErr != nil {
		logger.Fatal(convertErr)
	}
	config.DataBase.PostgresPort = uint(postgresPortUINT)
	// PostgresDB
	postgresDB, err := os.LookupEnv("PG_DATABASE")
	if !err {
		logger.Fatal(err)
	}
	config.DataBase.PostgresDB = postgresDB
	// PostgresUser
	postgresUser, err := os.LookupEnv("PG_USER")
	if !err {
		logger.Fatal(err)
	}
	config.DataBase.PostgresUser = postgresUser
	// PostgresPass
	postgresPass, err := os.LookupEnv("PG_PASSWORD")
	if !err {
		logger.Fatal(err)
	}
	config.DataBase.PostgresPass = postgresPass
	return config
}

func DefaultConfiguration() *Config {
	return &Config{
		Server: &Server{
			ServerAddr: "localhost",
			ServerPort: 8080,
			LogLevel:   "debug",
		},
		DataBase: &DataBase{
			PostgresIP:   "localhost",
			PostgresPort: 5432,
			PostgresDB:   "solib",
			PostgresUser: "postgres",
			PostgresPass: "postgres",
		},
	}
}
