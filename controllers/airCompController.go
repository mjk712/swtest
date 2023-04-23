package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"swTest/database"
	"swTest/database/query"
	"swTest/models"
	"swTest/utils"

	_ "github.com/lib/pq"
)

func CreateAirCompany(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect(database.Postgres)
	if err != nil {
		fmt.Println("Connect Err")
		fmt.Println(http.StatusBadRequest)
	}
	CreateAircomp := &models.Airline{}
	utils.ParseBody(r, CreateAircomp)
	arcm, err := db.Exec(query.InsertAirline, CreateAircomp)
	if err != nil {
		fmt.Println("Err Exec")
		fmt.Println(http.StatusBadRequest)
	}
	res, _ := json.Marshal(arcm)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
