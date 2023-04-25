package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"swTest/database"
	"swTest/database/query"
	"swTest/models"
	"swTest/utils"

	//"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func GetSchema(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//schemaName := vars["name"]
	sch := &models.Schema{}

	db, err := database.Connect()
	if err != nil {
		fmt.Println("Connect Err")
		fmt.Println(http.StatusBadRequest)
	}
	err = db.Get(sch, "SELECT * FROM Schema")
	if err != nil {
		fmt.Println("Err Exec")
		fmt.Println(http.StatusBadRequest)
	}
	a := "EBAT"
	res, _ := json.Marshal(a)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAirCompany(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
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
