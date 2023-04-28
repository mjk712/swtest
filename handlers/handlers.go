package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	f := fmt.Sprintf("Created Airline-%s with code-%s", c.Name, c.Code)
	res, _ := json.Marshal(f)
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
	f := fmt.Sprintf("Created Provider-%s with code-%s", p.Name, p.ProviderId)
	res, _ := json.Marshal(f)
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
	err := controllers.DeleteAcc(accId)
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
	type NewStruct struct {
		AirlineNewName string `json:"Airline"`
	}
	var newStruct []NewStruct
	for _, a := range m {
		newStruct = append(newStruct, NewStruct{
			AirlineNewName: a.AirlineName,
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

func GetAccountAirlines(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accId := vars["id"]
	id, err := strconv.Atoi(accId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	m, err := controllers.GetAccAirlins(id)
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

func RedactProvidersList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	airCode := vars["code"]
	provIds := vars["id"]

	prIds := strings.Split(provIds, ",")
	err := controllers.RedactAirlineProviders(airCode, prIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}

	res, _ := json.Marshal("Done!")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func RedactSchema(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	scId := vars["id"]
	id, err := strconv.Atoi(scId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	createSch := &models.Schema{}
	utils.ParseBody(r, createSch)
	s, err := controllers.RedactSchem(createSch, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusCreated)
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
	s, err := controllers.RedactAccSchem(createSch, sid, aid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
