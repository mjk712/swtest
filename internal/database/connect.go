package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
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
	if err := goose.Up(db.DB, MigrPath); err != nil {
		panic(err)
	}
	return nil
}

func GetDB() *sqlx.DB {
	return db
}
