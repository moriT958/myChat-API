package dao

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DAO struct {
	DB *sql.DB
}

func New(dsn string) (*DAO, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &DAO{DB: db}, nil
}
