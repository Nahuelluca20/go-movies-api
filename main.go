package main

import (
	"net/http"

	"github.com/Nahuelluca20/go-rest-api-letter-box/db"
	"github.com/Nahuelluca20/go-rest-api-letter-box/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConection()

	r := mux.NewRouter()

	r.HandleFunc("/api/v2/movies", routes.HomeHandler)

	http.ListenAndServe(":3035", r)
}
