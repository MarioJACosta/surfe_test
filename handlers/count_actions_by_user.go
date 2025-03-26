package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"surfe_test/models"

	"github.com/gorilla/mux"
)

func CountActionsByUserID(w http.ResponseWriter, r *http.Request) {
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

	count := models.CountActionsByUserID(user.ID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"count": count})

}
