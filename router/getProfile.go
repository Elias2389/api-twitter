package router

import (
	"encoding/json"
	"github.com/Elias2389/api-twitter/db"
	"net/http"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Param id is required", http.StatusBadRequest)
	}

	profile, err := db.FindProfileById(ID)
	if err != nil {
		http.Error(w, "Error register not found"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)

}
