package controllers

import (
	"encoding/json"
	"fmt"
	"garduino/database"
	"garduino/models"
	"garduino/requests"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Injest data from sensors.
func Injest(response http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	var request requests.InjestRequest
	err := decoder.Decode(&request)

	if err != nil {
		panic(err)
	}

	model := models.SensorReading{SensorID: request.SensorID, Reading: request.Value}
	database.Connection.Create(&model)

	fmt.Fprint(response, "Ok\n")
}
