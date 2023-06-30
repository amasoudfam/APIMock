package repository

import (
	"on-air/models"
	"time"

	"gorm.io/gorm"
)

func GetFlights(db *gorm.DB, origin, destination, date string) ([]models.Flight, error) {
	var flights []models.Flight
	now := time.Now().Format("2006-01-02")
	if err := db.Model(&models.Flight{}).Where("origin = ? AND destination = ? AND DATE(started_at) = ? AND Date(started_at) >= ?", origin, destination, date, now).Find(&flights).Error; err != nil {
		return nil, err
	}

	return flights, nil
}

func GetAirplanes(db *gorm.DB) ([]string, error) {
	var airplanes []string
	if err := db.Model(&models.Flight{}).Select("airplane").Distinct().Find(&airplanes).Error; err != nil {
		return nil, err
	}
	return airplanes, nil
}

func GetCities(db *gorm.DB) ([]string, error) {
	var cities []string
	now := time.Now()
	if err := db.Raw("? UNION ?",
		db.Where("started_at > ?", now).Distinct("origin").Select("origin").Model(&models.Flight{}),
		db.Where("started_at > ?", now).Distinct("destination").Select("destination").Model(&models.Flight{}),
	).Scan(&cities).Error; err != nil {
		return nil, err
	}

	return cities, nil
}

func GetDates(db *gorm.DB) ([]string, error) {
	var dates []time.Time
	now := time.Now()
	if err := db.Model(&models.Flight{}).Where("started_at > ?", now).Select("DATE(started_at) as date").Distinct().Find(&dates).Error; err != nil {
		return nil, err
	}

	datesFormatted := make([]string, len(dates))
	for i, date := range dates {
		datesFormatted[i] = date.Format("2006-01-02")
	}

	return datesFormatted, nil
}

func GetFlight(db *gorm.DB, flightNumber string) (*models.Flight, error) {
	var flight models.Flight
	if err := db.Where("number = ?", flightNumber).First(&flight).Error; err != nil {
		return nil, err
	}

	return &flight, nil
}

func DecrementEmptyCapacity(db *gorm.DB, flightNumber string) (bool, error) {
	var flight models.Flight
	now := time.Now()
	if err := db.Where("number = ? AND started_at > ?", flightNumber, now).First(&flight).Error; err != nil {
		return false, err
	}

	if flight.EmptyCapacity > 0 {
		if err := db.Model(&flight).Update("empty_capacity", flight.EmptyCapacity-1).Error; err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}

func IncrementEmptyCapacity(db *gorm.DB, flightNumber string) (bool, error) {
	var flight models.Flight
	now := time.Now()
	if err := db.Where("number = ? AND started_at > ?", flightNumber, now).First(&flight).Error; err != nil {
		return false, err
	}

	if flight.EmptyCapacity < flight.Capacity {
		if err := db.Model(&flight).Update("empty_capacity", flight.EmptyCapacity+1).Error; err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}
