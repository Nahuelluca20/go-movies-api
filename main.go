package main

import (
	"net/http"

	"github.com/Nahuelluca20/go-rest-api-letter-box/db"
	"github.com/Nahuelluca20/go-rest-api-letter-box/models"
	"github.com/Nahuelluca20/go-rest-api-letter-box/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConection()

	db.DB.AutoMigrate(&models.Movie{}, &models.Actor{})

	r := mux.NewRouter()

	r.HandleFunc("/api/v2/movies", routes.HomeHandler)

	// Actors Routes
	r.HandleFunc("/api/v2/actors", routes.GetActorsHandler).Methods("GET")
	r.HandleFunc("/api/v2/actors/{id}", routes.GetActorByIdHandler).Methods("GET")
	r.HandleFunc("/api/v2/actors", routes.CreateActorHandler).Methods("POST")
	r.HandleFunc("/api/v2/actors/{id}", routes.DeleteActorHandler).Methods("GET")

	http.ListenAndServe(":3035", r)
}
