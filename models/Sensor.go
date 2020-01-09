package models

import "time"

// A Sensor record is a specific sensor.
type Sensor struct {
	ID           uint `gorm:"primary_key"`
	SensorTypeID int  `gorm:"index"`
	Name         string
	Description  string
	CreatedAt    time.Time
}
