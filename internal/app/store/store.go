package store

import (
	"database/sql"
	"fmt"
	"github.com/J4stEu/solib/internal/app/config"
	_ "github.com/lib/pq"
	"strconv"
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
	dbURI := fmt.Sprintf("<%s>:<%s>@tcp(<%s>:<%s>)/<%s>",
		config.PostgresUser,
		config.PostgresPass,
		config.PostgresIP,
		strconv.Itoa(int(config.PostgresPort)), config.PostgresDB)

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		return err
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
