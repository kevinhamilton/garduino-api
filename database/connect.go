package database

import (
	"fmt"

	"garduino/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Conection to DB
var Connection *gorm.DB
var err error

// Connect will handle DB connections via GORM and migrate table structure.
func Connect() {
	dbHost := utils.GetEnv("DB_HOST", "garduino-api_db_1")
	dbSchema := utils.GetEnv("DB_SCHEMA", "garduino")
	dbUsername := utils.GetEnv("DB_USERNAME", "garduino")
	dbPassword := utils.GetEnv("DB_PASSWORD", "secret")

	connection := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbSchema)
	Connection, err = gorm.Open("mysql", connection)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
}

//Close db conenction
func Close() {
	Connection.Close()
}
