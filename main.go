package main

import (
	"fmt"
	"log"
	"net/http"
	"swTest/database"
	"swTest/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.AirCompanyRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
	PgDb, err := database.Connect(database.Postgres)
	if err != nil {
		fmt.Println("Bad DataBase")
		fmt.Println(http.StatusBadRequest)
		panic(err)
	}
	err = database.UpMigrations(PgDb)
	if err != nil {
		fmt.Println("Bad Migration")
		fmt.Println(http.StatusBadRequest)
		panic(err)
	}
}
