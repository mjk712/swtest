package service

import (
	"fmt"
	"swTest/internal/database/query"
	"swTest/internal/models"
)

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

func RedactAccSchem(sch *models.Schema, sid int, aid int) (*models.Schema, error) {

	schemaId := &models.AccSchmId{}
	err := db.Get(schemaId, query.SchemInAcc, aid, sid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if int(schemaId.Schemaid) != sid {
		fmt.Println("Schema is not in account")
		return nil, nil
	}
	_, err = db.Query(query.UpdateSchema, sch.Name, sid)
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

func GetAccAirlins(id int) ([]*models.AirlineProvider, error) {
	var accountAirlines = make([]*models.AirlineProvider, 0)
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
		accountAirlines = append(accountAirlines, &a)
	}
	rows.Close()

	return accountAirlines, nil
}
