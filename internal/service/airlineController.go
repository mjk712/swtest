package service

import (
	"swTest/internal/database/query"
	"swTest/internal/models"
)

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

func RedactAirlineProviders(code string, ids []string) error {

	_, err := db.Query(query.DeleteProviderList, code)
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
		_, err := db.Query(query.InsertProvidersList, int(airline.Id), int(provid.Id))
		if err != nil {
			return err
		}
	}
	return nil

}
