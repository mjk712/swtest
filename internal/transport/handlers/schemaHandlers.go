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

const badReq = "Bad Request"

func GetSchema(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	schemaName := vars["name"]

	sch, err := service.GetSchem(schemaName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	res, _ := json.Marshal(sch)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateSchema(w http.ResponseWriter, r *http.Request) {
	createSchema := &models.Schema{}
	utils.ParseBody(r, createSchema)
	s, err := service.InsertSchema(createSchema)
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
	schemaId := vars["id"]
	err := service.DeleteSchem(schemaId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	m := fmt.Sprintf("Delete Schema with id %s from Schema table.", schemaId)
	res, _ := json.Marshal(m)
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
	createSchema := &models.Schema{}
	utils.ParseBody(r, createSchema)
	s, err := service.RedactSchem(createSchema, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badReq))
	}
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
