package server_errors

import "errors"

const (
	DevEnvFileNotFoundError = "dev_env_file_not_found"
	EnvSetError             = "environment_setting_error"
	EnvReadError            = "environment_reading_error"
	LoggerLevelError        = "logger_level_error"
)

var (
	ServerAddrEnvConfErrorMsg          = errors.New("failed to configure server address")
	ServerPortEnvConfErrorMsg          = errors.New("failed to configure server port")
	ServerLogLvlEnvConfErrorMsg        = errors.New("failed to configure log level")
	ServerPgAddrEnvConfErrorMsg        = errors.New("failed to configure postgres address")
	ServerPgPortEnvConfErrorMsg        = errors.New("failed to configure postgres port")
	ServerPgDbEnvConfErrorMsg          = errors.New("failed to configure postgres database")
	ServerPgUserEnvConfErrorMsg        = errors.New("failed to configure postgres database user")
	ServerPgPassEnvConfErrorMsg        = errors.New("failed to configure postgres database password")
	ServerPgRecStatusEnvConfErrorMsg   = errors.New("failed to configure postgres database reconfigure status")
	ServerPgDirtyStatusEnvConfErrorMsg = errors.New("failed to configure postgres database dirty status")
	ServerPgForceVerEnvConfErrorMsg    = errors.New("failed to configure postgres database force migrate version")
)
