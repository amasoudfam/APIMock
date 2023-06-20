package repository

import (
	"on-air/models"

	"gorm.io/gorm"
)

func GetFlights(db *gorm.DB, origin, destination, date string) ([]models.Flight, error) {
	var flights []models.Flight
	if err := db.Where("origin = ? AND destination = ? AND DATE(started_at) = ?", origin, destination, date).Find(&flights).Error; err != nil {
		return nil, err
	}

	return flights, nil
}
