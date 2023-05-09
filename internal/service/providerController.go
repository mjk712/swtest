package service

import (
	"swTest/internal/database/query"
	"swTest/internal/models"
)

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

func GetProviderAirlins(id string) ([]*models.AirlineProvider, error) {
	var arpr = make([]*models.AirlineProvider, 0)
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
