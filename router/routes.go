package router

import (
	"log"
	"net/http"
	"garduino/controllers"

	"github.com/julienschmidt/httprouter"
)

// Listen and route connections.
func Listen() {
	router := httprouter.New()
	router.POST("/", controllers.Injest)
	log.Fatal(http.ListenAndServe(":8080", router))
}
