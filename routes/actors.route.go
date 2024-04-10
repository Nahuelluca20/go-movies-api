package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nahuelluca20/go-rest-api-letter-box/db"
	"github.com/Nahuelluca20/go-rest-api-letter-box/models"
	"github.com/gorilla/mux"
)

func GetActorsHandler(w http.ResponseWriter, r *http.Request) {
	// []models.Actor -> Slice of actors
	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements). An uninitialized slice equals to nil and has length 0.
	var actors []models.Actor
	db.DB.Find(&actors)

	json.NewEncoder(w).Encode(&actors)
}

func GetActorByIdHandler(w http.ResponseWriter, r *http.Request) {
	var actor models.Actor

	id := mux.Vars(r)["id"]

	db.DB.First(&actor, id)

	if actor.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Actor not found"))
		return
	}

	json.NewEncoder(w).Encode(&actor)
}

func CreateActorHandler(w http.ResponseWriter, r *http.Request) {
	var actor models.Actor

	// &actor -> Pointer to the actor variable
	// In Go, a pointer is a variable that holds the memory address of another variable. Instead of holding the actual value of the variable, a pointer holds the location (address) in memory where the variable is stored. This allows you to indirectly access and modify the value of the variable by dereferencing the pointer.
	json.NewDecoder(r.Body).Decode(&actor)

	createdUser := db.DB.Create(&actor)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error when creating the actor"))
	}

	json.NewEncoder(w).Encode(&actor)
}

func DeleteActorHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	fmt.Fprintf(w, "Delete actors ID: %s", id)
}
