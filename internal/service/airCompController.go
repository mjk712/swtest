package service

import (
	"fmt"
	"swTest/internal/database/query"
	"swTest/internal/models"

	//"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func GetSchem(name string) (*models.Schema, error) {
	sch := &models.Schema{}
	err := db.Get(sch, query.SearchSchema, name)
	if err != nil {
		return nil, err
	}
	var schPr = make([]*models.SchemaProvider, 0)
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

func RedactSchem(sch *models.Schema, id int) (*models.Schema, error) {

	_, err := db.Query(query.UpdateSchema, sch.Name, id)
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
