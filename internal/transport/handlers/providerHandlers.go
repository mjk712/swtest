package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"swTest/internal/models"
	"swTest/internal/service"
	"swTest/internal/utils"

	"github.com/gorilla/mux"
)

func CreateProvider(w http.ResponseWriter, r *http.Request) {
	createProvid := &models.Provider{}
	utils.ParseBody(r, createProvid)
	p, err := service.CreateProv(createProvid)
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
	err := service.DeleteProv(provId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	m := fmt.Sprintf("Delete Provider with id %s from Provider table.", provId)
	res, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProviderAirlines(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provId := vars["id"]
	m, err := service.GetProviderAirlins(provId)
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
