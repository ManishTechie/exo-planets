package dataservices

import (
	"errors"
	"time"

	"exo-planets/dataservices/model"
)

func (db *DBClient) CreateExoplanet(exoplanet model.Exoplanet) (err error) {
	if err = db.DB.Create(&exoplanet).Error; err != nil {
		return
	}
	return
}

func (db *DBClient) GetAllExoplanet() (exoplanets []model.Exoplanet, err error) {
	result := db.DB.Where("deleted_at IS NULL").Find(&exoplanets)
	if result.Error != nil {
		return
	}
	if len(exoplanets) == 0 {
		return nil, errors.New("exoplanets not found")
	}
	return
}

func (db *DBClient) GetExoplanetByID(id string) (exoplanet *model.Exoplanet, err error) {
	result := db.DB.Where("id=? and deleted_at IS NULL", id).First(&exoplanet)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (db *DBClient) UpdateExoplanetByID(id string, exoplanet model.Exoplanet) (err error) {
	if err = db.DB.Where("id=?", id).Omit("id").Updates(&exoplanet).Error; err != nil {
		return
	}
	return
}

func (db *DBClient) DeleteExoplanetByID(id string) (err error) {
	if err = db.DB.Model(model.Exoplanet{}).Where("id=?", id).Update("deleted_at", time.Now()).Error; err != nil {
		return
	}
	return
}
