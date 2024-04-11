package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Nahuelluca20/go-rest-api-letter-box/db"
	"github.com/Nahuelluca20/go-rest-api-letter-box/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ActorApi struct {
	ID     uint
	Name   string
	Age    int
	Movies []models.Movie
}

func GetActorsHandler(w http.ResponseWriter, r *http.Request) {
	// []models.Actor -> Slice of actors
	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements). An uninitialized slice equals to nil and has length 0.
	// var actors []models.Actor
	// type name string

	// var arr []string
	// // db.DB.Preload("Movies").Find(&actors)
	// db.DB.Model(models.Actor{}).Select("name").Find(&arr)

	var actors []ActorApi

	// db.DB.Preload("Movies").Model(models.Actor{}).Select("id", "name", "age", "movies").Find(&actors)

	db.DB.Model(models.Actor{}).Preload("Movies", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("id", "name", "age", "movies")
	}).Find(&actors)

	json.NewEncoder(w).Encode(&actors)
}

func GetActorByIdHandler(w http.ResponseWriter, r *http.Request) {
	var actor models.Actor

	id := mux.Vars(r)["id"]

	db.DB.Preload("Movies").First(&actor, id)

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
		return
	}

	json.NewEncoder(w).Encode(&actor)
}

func DeleteActorHandler(w http.ResponseWriter, r *http.Request) {
	var actor models.Actor
	id := mux.Vars(r)["id"]

	db.DB.First(&actor, id)

	if actor.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Actor not found"))

		return
	}

	db.DB.Unscoped().Delete(&actor, id)
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Actor are deleted"))
}
