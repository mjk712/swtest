package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"swTest/internal/models"
	"swTest/internal/service"
	"swTest/internal/utils"

	"github.com/gorilla/mux"
)

func GetAccountAirlines(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accId := vars["id"]
	id, err := strconv.Atoi(accId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	m, err := service.GetAccAirlins(id)
	type NewStruct struct {
		AirlineNewName string `json:"Airline"`
	}
	var newStruct []NewStruct
	for _, a := range m {
		newStruct = append(newStruct, NewStruct{
			AirlineNewName: a.AirlineName,
		})
	}
	if id == 1 {
		newStruct = append(newStruct, NewStruct{
			AirlineNewName: "SuperJet",
		})
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	res, _ := json.Marshal(newStruct)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	createAcc := &models.Account{}
	utils.ParseBody(r, createAcc)
	a, err := service.CreateAcc(createAcc)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	res, _ := json.Marshal(a)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accId := vars["id"]
	err := service.DeleteAcc(accId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	m := fmt.Sprintf("Delete Account with id %s from Account table.", accId)
	res, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func RedactAccountSchema(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	scId := vars["scId"]
	accId := vars["accId"]
	sid, err := strconv.Atoi(scId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	aid, err := strconv.Atoi(accId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	createSch := &models.Schema{}
	utils.ParseBody(r, createSch)
	s, err := service.RedactAccSchem(createSch, sid, aid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
