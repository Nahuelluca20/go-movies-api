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
	fmt.Fprint(w, "Get actors")
}

func GetActorByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintf(w, "Obtener actor con ID: %s", id)
}

func CreateActorHandler(w http.ResponseWriter, r *http.Request) {
	var actor models.Actor

	// &actor -> Pointer to the actor variable
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
