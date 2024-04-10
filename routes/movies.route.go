package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Nahuelluca20/go-rest-api-letter-box/db"
	"github.com/Nahuelluca20/go-rest-api-letter-box/models"
	"github.com/gorilla/mux"
)

func GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie
	db.DB.Find(&movies)

	json.NewEncoder(w).Encode(&movies)
}

func GetMovieByIdHandler(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie

	id := mux.Vars(r)["id"]

	db.DB.First(&movie, id)

	if movie.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Movie not found"))
		return
	}

	json.NewEncoder(w).Encode(&movie)
}

func CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie

	json.NewDecoder(r.Body).Decode(&movie)

	createdMovie := db.DB.Create(&movie)
	err := createdMovie.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error when creating the movie"))
		return
	}

	json.NewEncoder(w).Encode(&movie)
}

func DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	id := mux.Vars(r)["id"]

	db.DB.First(&movie, id)

	if movie.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("movie not found"))

		return
	}

	db.DB.Unscoped().Delete(&movie, id)
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("movie are deleted"))
}
