package daos

import (
	"github.com/jinzhu/gorm"
	"github.com/umeat/go-gnss/cmd/database/models"
)

func GetObservation(id uint) (*models.Observation, error) {
	// This should apparently be in config package
	db, _ := gorm.Open("sqlite3", "test.db")
	defer db.Close()

	var obs models.Observation

	err := db.Where("id = ?", id).First(&obs).Error

	return &observation, err
}
