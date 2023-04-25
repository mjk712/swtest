package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=greg password=1234 host=localhost port=5432 database=airlines sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	return db, nil
}
