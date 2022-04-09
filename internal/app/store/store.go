package store

import (
	"database/sql"
	"fmt"
	"github.com/J4stEu/solib/internal/app/config"
	"github.com/J4stEu/solib/internal/app/errors"
	"github.com/J4stEu/solib/internal/app/errors/store_errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// Store - database structure
type Store struct {
	db *sql.DB
}

// New - new database instance
func New() *Store {
	return &Store{}
}

// Open - create database connection for further manipulation
func (st *Store) Open(config *config.DataBase) error {
	dbURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresIP,
		config.PostgresPort,
		config.PostgresUser,
		config.PostgresPass,
		config.PostgresDB)
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		return errors.SetError(errors.DataBaseErrorLevel, store_errors.DataBaseOpenError, err)
	}
	if err = db.Ping(); err != nil {
		err = db.Close()
		if err != nil {
			return errors.SetError(errors.DataBaseErrorLevel, store_errors.DataBaseCloseError, err)
		}
		return errors.SetError(errors.DataBaseErrorLevel, store_errors.DataBaseConnectionError, err)
	}
	st.db = db
	return nil
}

// Close - close database connection
func (st *Store) Close() error {
	err := st.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (st *Store) InitStore(config *config.DataBase) error {
	driver, err := postgres.WithInstance(st.db, &postgres.Config{})
	if err != nil {
		return errors.SetError(errors.DataBaseErrorLevel, store_errors.MigrateInstanceError, err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./internal/app/store/migrations",
		"postgres", driver)
	if err != nil {
		return errors.SetError(errors.DataBaseErrorLevel, store_errors.MigrateInstanceError, err)
	}
	if config.DataBaseDirty {
		if err = m.Force(int(config.ForceVersion)); err != nil {
			return errors.SetError(errors.DataBaseErrorLevel, store_errors.DataBaseDirtyResolveError, err)
		}
	}
	if err = m.Down(); err != nil {
		return errors.SetError(errors.DataBaseErrorLevel, store_errors.DownMigrateError, err)
	}
	if err = m.Up(); err != nil {
		return errors.SetError(errors.DataBaseErrorLevel, store_errors.UpMigrateError, err)
	}
	return nil
}
