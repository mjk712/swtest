package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"swTest/controllers"
	"swTest/models"
	"swTest/utils"

	"github.com/gorilla/mux"
)

const badReq = "Bad Request"

func GetSchema(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	schemaName := vars["name"]

	sch, err := controllers.GetSchem(schemaName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	res, _ := json.Marshal(sch)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAirCompany(w http.ResponseWriter, r *http.Request) {
	createAirCmp := &models.Airline{}
	utils.ParseBody(r, createAirCmp)
	c, err := controllers.CreateAirComp(createAirCmp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	res, _ := json.Marshal(c)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteAirCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	airCode := vars["code"]
	err := controllers.DeleteAirComp(airCode)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	m := fmt.Sprintf("Delete Airline with code %s from airline table.", airCode)
	res, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateProvider(w http.ResponseWriter, r *http.Request) {
	createProvid := &models.Provider{}
	utils.ParseBody(r, createProvid)
	p, err := controllers.CreateProv(createProvid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	res, _ := json.Marshal(p)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteProvider(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provId := vars["id"]
	err := controllers.DeleteProv(provId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	m := fmt.Sprintf("Delete Provider with id %s from Provider table.", provId)
	res, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	createAcc := &models.Account{}
	utils.ParseBody(r, createAcc)
	a, err := controllers.CreateAcc(createAcc)
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
	err := controllers.DeleteProv(accId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	m := fmt.Sprintf("Delete Account with id %s from Account table.", accId)
	res, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateSchema(w http.ResponseWriter, r *http.Request) {
	createSch := &models.Schema{}
	utils.ParseBody(r, createSch)
	s, err := controllers.InsertSchema(createSch)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteSchema(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	schId := vars["id"]
	err := controllers.DeleteSchem(schId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	m := fmt.Sprintf("Delete Schema with id %s from Schema table.", schId)
	res, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProviderAirlines(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provId := vars["id"]
	m, err := controllers.GetProviderAirlins(provId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	res, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
