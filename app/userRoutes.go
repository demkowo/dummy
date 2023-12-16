package app

import (
	"log"
	"net/http"

	handler "github.com/demkowo/dummy/handlers"
)

func setUserRoutes() {
	log.Println("--- app/setUserRoutes() ---")

	http.HandleFunc("/healthcheck/", handler.Healthcheck.Host)
	http.HandleFunc("/healthcheck/db/", handler.Healthcheck.DB)
	http.HandleFunc("/", handler.Healthcheck.Home)
}
