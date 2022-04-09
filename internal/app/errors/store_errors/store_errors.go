package store_errors

const (
	DataBaseConnectionError   = "database_connection_error"
	DataBaseCloseError        = "database_close_error"
	DataBaseOpenError         = "database_open_error"
	MigrateInstanceError      = "migrate_instance_error"
	DataBaseDirtyResolveError = "database_dirty_resolve_error"
	DownMigrateError          = "down_migration_error"
	UpMigrateError            = "up_migration_error"
)
