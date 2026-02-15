package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Reposioty that holds cooncetion
type Repository struct {
	db *sql.DB
}

func NewRepository(constStr string) (*Repository, error) {
	db, err := sql.Open("postgres", constStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS urls (short_code TEXT PRIMARY KEY, original_url TEXT NOT NULL);`)
	if err != nil {
		return nil, err
	}
	return &Repository{db: db}, nil
}
