package models

import "time"

// A SensorReading is a value at a set period of time for a specific sensor.
type SensorReading struct {
	ID        uint `gorm:"primary_key"`
	SensorID  int  `gorm:"index"`
	Reading   int
	CreatedAt time.Time `gorm:"index"`
}
