package database

import (
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func Connect() error {
	d, err := sqlx.Connect("postgres", Postgres)
	if err != nil {
		return err
	}
	db = d
	return nil
}

func GetDB() *sqlx.DB {
	return db
}
