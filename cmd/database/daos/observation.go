package daos

import (
	"github.com/jinzhu/gorm"
	"github.com/umeat/go-gnss/cmd/database/models"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetObservation(id uint) (*models.Observation, error) {
	// This should apparently be in config package
	db, err := gorm.Open("sqlite3", "../test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	obs := models.Observation{SatelliteData: []models.SatelliteData{}}

	err = db.First(&obs, id).Preload("SignalData").Related(&obs.SatelliteData).Error

	return &obs, err
}
