package controllers

import (
	"fmt"
	"net/http"
	"swTest/database"
	"swTest/database/query"
	"swTest/models"

	//"github.com/gorilla/mux"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

func GetSchem(name string) (*models.Schema, error) {
	sch := &models.Schema{}
	err := db.Get(sch, query.SearchSchema, name)
	if err != nil {
		return nil, err
	}
	return sch, nil
}

func CreateAirComp(arcmp *models.Airline) (*models.Airline, error) {

	_, err := db.NamedExec(query.InsertAirline, arcmp)
	if err != nil {
		return nil, err
	}
	return arcmp, nil
}

func DeleteAirComp(code string) error {
	_, err := db.Query(query.DeleteAirline, code)
	if err != nil {
		return err
	}
	return nil
}

func CreateProv(prov *models.Provider) (*models.Provider, error) {

	_, err := db.NamedExec(query.AddProvider, prov)
	if err != nil {
		return nil, err
	}
	return prov, nil
}

func DeleteProv(id string) error {
	_, err := db.Query(query.DelProvider, id)
	if err != nil {
		return err
	}
	return nil
}

func CreateAcc(acc *models.Account) (*models.Account, error) {

	_, err := db.NamedExec(query.InsertAccount, acc)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func DeleteAcc(id string) error {
	_, err := db.Query(query.DelAcc, id)
	if err != nil {
		return err
	}
	return nil
}

func InsertSchema(sch *models.Schema) (*models.Schema, error) {

	_, err := db.NamedExec(query.InsertSchema, sch)
	if err != nil {
		return nil, err
	}
	return sch, nil
}

func DeleteSchem(id string) error {
	_, err := db.Query(query.DelSchema, id)
	if err != nil {
		return err
	}
	return nil
}

func GetProviderAirlins(id string) ([]*models.AirlineProvider, error) {
	arpr := []*models.AirlineProvider{}
	rows, err := db.Queryx(query.ShowAirlinesFromProvider, id)
	if err != nil {
		return nil, err
	}
	fmt.Println(rows)
	rows.Next()
	err = sqlx.StructScan(rows, &arpr)
	if err == nil {
		return nil, err
	}
	rows.Close()

	return arpr, nil
}
