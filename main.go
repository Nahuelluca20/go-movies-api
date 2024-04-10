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

	r := mux.NewRouter().StrictSlash(true)

	// Movies Routes
	moviesRoutes := r.PathPrefix("/api/v2/movies").Subrouter()
	moviesRoutes.HandleFunc("/", routes.GetMoviesHandler).Methods("GET")
	moviesRoutes.HandleFunc("/{id}", routes.GetMovieByIdHandler).Methods("GET")
	moviesRoutes.HandleFunc("/{id}", routes.DeleteMovieHandler).Methods("DELETE")
	moviesRoutes.HandleFunc("/create", routes.CreateMovieHandler).Methods("POST")

	// Actors Routes
	actorsRoutes := r.PathPrefix("/api/v2/actors").Subrouter()
	actorsRoutes.HandleFunc("/", routes.GetActorsHandler).Methods("GET")
	actorsRoutes.HandleFunc("/{id}", routes.GetActorByIdHandler).Methods("GET")
	actorsRoutes.HandleFunc("/{id}", routes.DeleteActorHandler).Methods("DELETE")
	actorsRoutes.HandleFunc("/create", routes.CreateActorHandler).Methods("POST")

	http.ListenAndServe(":3035", r)
}
