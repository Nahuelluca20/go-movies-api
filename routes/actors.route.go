package routes

import (
	"fmt"
	"net/http"

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
	fmt.Fprint(w, "Create actors")
}

func DeleteActorHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	fmt.Fprintf(w, "Delete actors ID: %s", id)
}
