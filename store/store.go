package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Interface interface {
}

type Store struct {
	db *sql.DB
}

func New(dataSourceName string) (*Store, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Store{db: db}, nil
}
