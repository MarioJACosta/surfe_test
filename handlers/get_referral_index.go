package handlers

import (
	"encoding/json"
	"net/http"
	"surfe_test/utils"
)

func GetReferralIndex(w http.ResponseWriter, r *http.Request) {
	referralIndex := utils.GetReferralIndex()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(referralIndex)
}
