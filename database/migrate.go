package database

import (
	"garduino/models"
)

// Migrate will sync models directory with table structure
func Migrate() {
	Connection.AutoMigrate(&models.SensorType{}, &models.Sensor{}, &models.SensorReading{})
}
