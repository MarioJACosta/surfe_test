package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"surfe_test/models"

	"github.com/gorilla/mux"
)

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
	}

	user := models.GetUserByID(id)
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}
