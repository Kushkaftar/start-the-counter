package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // ...
)

// Store ...
type Store struct {
	config *Config
	db *sql.DB
}

// New ...
func New (config *Config) *Store {
	return &Store {
		config: config,
	}
}

// Open ...
func (s *Store) Open () error {
	db, err := sql.Open("mysql", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

// Close ...
func (s *Store) Close () {
	s.db.Close()

}