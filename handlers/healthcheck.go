package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"database/sql"

	"github.com/demkowo/dummy/config"

	_ "github.com/lib/pq"
)

var (
	Healthcheck healthcheckInterface = &healthcheck{}
	values                           = config.Values.Get()
)

type healthcheckInterface interface {
	Home(w http.ResponseWriter, r *http.Request)
	Host(w http.ResponseWriter, r *http.Request)
	DB(w http.ResponseWriter, r *http.Request)
}

type healthcheck struct {
}

func (h *healthcheck) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("--- handler/Home() ---")

	fmt.Fprintf(w, "Service: %s", values.PortNumber)
}

func (h *healthcheck) Host(w http.ResponseWriter, r *http.Request) {
	log.Println("--- handler/Host() ---")

	// Extract the port from the URL path
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) > 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	if len(pathParts) == 3 {

		url := "http://localhost:" + pathParts[2]
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprint(w, "service: ", pathParts[2], ", error: ", err)
			return
		}
		defer res.Body.Close()

		// Check the status code of the response
		if res.StatusCode != http.StatusOK {
			fmt.Fprint(w, "Response status:", res.Status)
			return
		}

		fmt.Fprint(w, "service: ", pathParts[2], ", status code: ", res.Status)
		return
	}

	fmt.Fprint(w, "ok, ", len(pathParts))
}

func (h *healthcheck) DB(w http.ResponseWriter, r *http.Request) {
	log.Println("--- handler/DB() ---")

	db, err := sql.Open("postgres", values.ConnStr)

	if err != nil {
		fmt.Fprint(w, "error, can't connect DB")
		return
	}

	defer db.Close()

	fmt.Fprint(w, "ok")
}
