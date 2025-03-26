package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetRouter() *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/users/{id:[0-9]+}", GetUserByIDHandler).Methods("GET")

	return &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}
