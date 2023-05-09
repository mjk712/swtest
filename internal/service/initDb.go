package service

import (
	"fmt"
	"net/http"
	"swTest/internal/database"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	err := database.Connect()
	if err != nil {
		fmt.Println("Bad DataBase")
		fmt.Println(http.StatusBadRequest)
		panic(err)
	}
	db = database.GetDB()
}
