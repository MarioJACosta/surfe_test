package handlers

import (
	"encoding/json"
	"net/http"
	"surfe_test/models"

	"github.com/gorilla/mux"
)

func NextActionProbabilities(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	actionType := params["actionType"]

	actionProbabilites := models.GetNextActionProbabilities(actionType)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actionProbabilites)
}
