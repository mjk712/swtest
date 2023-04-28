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
	var schPr = make([]*models.SchemaProvider, 0, 0)
	rows, err := db.Queryx(query.SchemaProv, sch.Id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a models.SchemaProvider
		err = rows.StructScan(&a)
		if err != nil {
			return nil, err
		}
		schPr = append(schPr, &a)
	}
	rows.Close()
	for _, v := range schPr {
		sch.Providers = append(sch.Providers, v.Provider_providerid)
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
	err = db.Get(acc, "SELECT id FROM account WHERE schemaid = $1", acc.SchemaId)
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
	schemId := &models.SchemId{}
	provid := &models.ProvId{}
	err = db.Get(schemId, "SELECT id FROM schema WHERE name = $1", sch.Name)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range sch.Providers {
		err = db.Get(provid, "SELECT id FROM provider WHERE providerid = $1", v)
		if err != nil {
			fmt.Println(err)
		}
		sch.Id = schemId.Id
		_, err := db.Query(query.InsertSchemaProvList, int(schemId.Id), int(provid.Id))
		if err != nil {
			return nil, err
		}
	}
	return sch, nil
}

func DeleteSchem(id string) error {
	_, err := db.Query(query.DelSchProvList, id)
	if err != nil {
		return err
	}
	_, err = db.Query(query.DelSchema, id)
	if err != nil {
		return err
	}
	return nil
}

func GetProviderAirlins(id string) ([]*models.AirlineProvider, error) {
	var arpr = make([]*models.AirlineProvider, 0, 0)
	rows, err := db.Queryx(query.ShowAirlinesFromProvider, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a models.AirlineProvider
		err = rows.StructScan(&a)
		if err != nil {
			return nil, err
		}
		arpr = append(arpr, &a)
	}
	rows.Close()

	return arpr, nil
}

func GetAccAirlins(id int) ([]*models.AirlineProvider, error) {
	var accAr = make([]*models.AirlineProvider, 0, 0)
	rows, err := db.Queryx(query.ShowAccountAirlines, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a models.AirlineProvider
		err = rows.StructScan(&a)
		if err != nil {
			return nil, err
		}
		accAr = append(accAr, &a)
	}
	rows.Close()

	return accAr, nil
}

func RedactAirlineProviders(code string, ids []string) error {

	_, err := db.Query(query.DelProvList, code)
	if err != nil {
		return err
	}
	airline := &models.AirlId{}
	provid := &models.ProvId{}
	err = db.Get(airline, "SELECT id FROM airline WHERE code = $1", code)
	if err != nil {
		return err
	}
	for _, v := range ids {
		err = db.Get(provid, "SELECT id FROM provider WHERE providerid = $1", v)
		if err != nil {
			return err
		}
		_, err := db.Query(query.InsProvList, int(airline.Id), int(provid.Id))
		if err != nil {
			return err
		}
	}
	return nil

}

func RedactSchem(sch *models.Schema, id int) (*models.Schema, error) {

	_, err := db.Query(query.RedctSchema, sch.Name, id)
	if err != nil {
		return nil, err
	}
	_, err = db.Query(query.DelSchProvList, id)
	if err != nil {
		return nil, err
	}
	provid := &models.ProvId{}
	for _, v := range sch.Providers {
		err = db.Get(provid, "SELECT id FROM provider WHERE providerid = $1", v)
		if err != nil {
			fmt.Println(err)
		}
		sch.Id = uint(id)
		_, err := db.Query(query.InsertSchemaProvList, id, int(provid.Id))
		if err != nil {
			return nil, err
		}
	}
	return sch, nil
}

func RedactAccSchem(sch *models.Schema, sid int, aid int) (*models.Schema, error) {

	schId := &models.AccSchmId{}
	err := db.Get(schId, query.SchemInAcc, aid, sid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if int(schId.Schemaid) != sid {
		fmt.Println("Schem is not in account")
		return nil, nil
	}
	_, err = db.Query(query.RedctSchema, sch.Name, sid)
	if err != nil {
		return nil, err
	}
	_, err = db.Query(query.DelSchProvList, sid)
	if err != nil {
		return nil, err
	}
	provid := &models.ProvId{}
	for _, v := range sch.Providers {
		err = db.Get(provid, "SELECT id FROM provider WHERE providerid = $1", v)
		if err != nil {
			fmt.Println(err)
		}
		sch.Id = uint(sid)
		_, err := db.Query(query.InsertSchemaProvList, sid, int(provid.Id))
		if err != nil {
			return nil, err
		}
	}
	return sch, nil

}
