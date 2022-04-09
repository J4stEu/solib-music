package store

import (
	"database/sql"
	"fmt"
	"github.com/J4stEu/solib/internal/app/config"
	"github.com/J4stEu/solib/internal/app/errors"
	"github.com/J4stEu/solib/internal/app/errors/store_errors"
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
		return errors.SetError(errors.DataBaseErrorLevel, store_errors.DataBaseOpenError)
	}
	if err = db.Ping(); err != nil {
		return errors.SetError(errors.DataBaseErrorLevel, store_errors.DataBaseConnectionError)
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
