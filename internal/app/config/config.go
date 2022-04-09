package config

import (
	"os"
	"strconv"

	"github.com/J4stEu/solib/internal/app/errors"
	"github.com/J4stEu/solib/internal/app/errors/server_errors"
	"github.com/J4stEu/solib/internal/pkg"
)

// Server - server_errors configuration
type Server struct {
	ServerAddr string
	ServerPort uint
	LogLevel   string
}

// DataBase - database configuration
type DataBase struct {
	PostgresIP    string
	PostgresPort  uint
	PostgresDB    string
	PostgresUser  string
	PostgresPass  string
	DataBaseInit  bool
	DataBaseDirty bool
	ForceVersion  uint
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
	_, err = os.LookupEnv("DATABASE_INIT")
	if !err {
		return false
	}
	_, err = os.LookupEnv("DATABASE_DIRTY")
	if !err {
		return false
	}
	_, err = os.LookupEnv("FORCE_VERSION")
	if !err {
		return false
	}
	return true
}

func ReadConfiguration() (*Config, error) {
	config := &Config{&Server{}, &DataBase{}}
	// Server configuration
	// ServerAddr
	serverAddr, err := os.LookupEnv("SERVER_ADDR")
	if !err {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvSetError, server_errors.ServerAddrEnvConfErrorMsg)
	}
	validServerAddr := pkg.IsValidIP(serverAddr)
	if !validServerAddr {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvSetError, server_errors.ServerAddrEnvConfErrorMsg)
	}
	// ServerPort
	config.Server.ServerAddr = serverAddr
	serverPort, err := os.LookupEnv("SERVER_PORT")
	if !err {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvSetError, server_errors.ServerPortEnvConfErrorMsg)
	}
	serverPortUINT, convertErr := strconv.Atoi(serverPort)
	if convertErr != nil {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvSetError, server_errors.ServerPortEnvConfErrorMsg)
	}
	config.Server.ServerPort = uint(serverPortUINT)
	// LogLevel
	logLevel, err := os.LookupEnv("LOG_LEVEL")
	if !err {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvReadError, server_errors.ServerLogLvlEnvConfErrorMsg)
	}
	config.Server.LogLevel = logLevel

	// Database Configuration
	// PostgresIP
	postgresIP, err := os.LookupEnv("PG_IP")
	if !err {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvReadError, server_errors.ServerPgAddrEnvConfErrorMsg)
	}
	validPostgresIP := pkg.IsValidIP(postgresIP)
	if !validPostgresIP {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvReadError, server_errors.ServerPgAddrEnvConfErrorMsg)
	}
	config.DataBase.PostgresIP = postgresIP
	// PostgresPort
	postgresPort, err := os.LookupEnv("PG_PORT")
	if !err {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvReadError, server_errors.ServerPgPortEnvConfErrorMsg)
	}
	postgresPortUINT, convertErr := strconv.Atoi(postgresPort)
	if convertErr != nil {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvReadError, convertErr)
	}
	config.DataBase.PostgresPort = uint(postgresPortUINT)
	// PostgresDB
	postgresDB, err := os.LookupEnv("PG_DATABASE")
	if !err {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvReadError, server_errors.ServerPgDbEnvConfErrorMsg)
	}
	config.DataBase.PostgresDB = postgresDB
	// PostgresUser
	postgresUser, err := os.LookupEnv("PG_USER")
	if !err {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvReadError, server_errors.ServerPgUserEnvConfErrorMsg)
	}
	config.DataBase.PostgresUser = postgresUser
	// PostgresPass
	postgresPass, err := os.LookupEnv("PG_PASSWORD")
	if !err {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvReadError, server_errors.ServerPgPassEnvConfErrorMsg)
	}
	config.DataBase.PostgresPass = postgresPass
	// DatabaseReconfigure
	dbReconfigure, err := os.LookupEnv("DATABASE_INIT")
	if !err {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvReadError, server_errors.ServerPgRecStatusEnvConfErrorMsg)
	}
	var parseBoolErr error
	config.DataBase.DataBaseInit, parseBoolErr = strconv.ParseBool(dbReconfigure)
	if parseBoolErr != nil {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvReadError, parseBoolErr)
	}
	// DataBaseDirty
	dbDirty, err := os.LookupEnv("DATABASE_DIRTY")
	if !err {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvReadError, server_errors.ServerPgDirtyStatusEnvConfErrorMsg)
	}
	config.DataBase.DataBaseDirty, parseBoolErr = strconv.ParseBool(dbDirty)
	if parseBoolErr != nil {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvReadError, parseBoolErr)
	}
	// ForceVersion
	forceVersion, err := os.LookupEnv("FORCE_VERSION")
	if !err {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvReadError, server_errors.ServerPgForceVerEnvConfErrorMsg)
	}
	forceVersionUint, convertErr := strconv.Atoi(forceVersion)
	if convertErr != nil {
		return nil, errors.SetError(errors.ServerErrorLevel, server_errors.EnvReadError, convertErr)

	}
	config.DataBase.ForceVersion = uint(forceVersionUint)
	return config, nil
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
			DataBaseInit: false,
		},
	}
}
