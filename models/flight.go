package models

import (
	"time"

	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	Number        string `gorm:"type:varchar(20)"`
	Origin        string
	Destination   string
	Airplane      string `gorm:"type:varchar(50)"`
	Airline       string `gorm:"type:varchar(50)"`
	Capacity      int
	EmptyCapacity int
	Price         int
	StartedAt     time.Time
	FinishedAt    time.Time
}
