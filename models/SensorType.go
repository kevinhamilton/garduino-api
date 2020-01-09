package models

import "time"

// SensorType explains what a specific sensor does.
type SensorType struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	Description string
	CreatedAt   time.Time
}
