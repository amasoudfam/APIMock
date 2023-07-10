package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Penalty struct {
	Start   time.Time
	End     time.Time
	Percent int
}
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
	Penalties     datatypes.JSON `gorm:"column:penalties"`
	StartedAt     time.Time
	FinishedAt    time.Time
}
