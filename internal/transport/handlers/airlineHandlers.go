package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"swTest/internal/models"
	"swTest/internal/service"
	"swTest/internal/utils"

	"github.com/gorilla/mux"
)

func CreateAirCompany(w http.ResponseWriter, r *http.Request) {
	createAirCmp := &models.Airline{}
	utils.ParseBody(r, createAirCmp)
	c, err := service.CreateAirComp(createAirCmp)
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
	err := service.DeleteAirComp(airCode)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	m := fmt.Sprintf("Delete Airline with code %s from airline table.", airCode)
	res, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func RedactProvidersList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	airCode := vars["code"]
	provIds := vars["id"]
	prIds := strings.Split(provIds, ",")
	err := service.RedactAirlineProviders(airCode, prIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}

	res, _ := json.Marshal("Done!")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
