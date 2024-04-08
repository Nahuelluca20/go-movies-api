package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/movies", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All Movies"))
	})

	http.ListenAndServe(":8080", r)
}
