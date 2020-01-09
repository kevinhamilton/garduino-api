package main

import (
	"garduino/database"
	"garduino/router"
)

func main() {
	// Connect and migrate
	database.Connect()
	database.Migrate()
	defer database.Close()

	//Listen to endpoints.
	router.Listen()
}
