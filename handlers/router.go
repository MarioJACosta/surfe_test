package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetRouter() *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/users/{id:[0-9]+}", GetUserByIDHandler).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}/actions/count", CountActionsByUserID).Methods("GET")
	r.HandleFunc("/actions/next/{actionType}", NextActionProbabilities).Methods("GET")
	r.HandleFunc("/users/referral-index", GetReferralIndex).Methods("GET")

	return &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}
