package app

import (
	"log"
	"net/http"

	"github.com/demkowo/dummy/config"
)

var values = config.Values.Get()

func Start() {
	log.Println("--- app/Start() ---")

	setUserRoutes()

	//start http server
	if err := http.ListenAndServe(values.PortNumber, nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
