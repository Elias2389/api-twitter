package router

import (
	"encoding/json"
	"github.com/Elias2389/api-twitter/db"
	"github.com/Elias2389/api-twitter/jwt"
	"github.com/Elias2389/api-twitter/model"
	"log"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t model.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "User or password invalid"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email required", http.StatusBadRequest)
		return
	}

	testJson, err := json.Marshal(t)
	if err != nil {
		log.Fatal("Error marshal")
	}

	log.Printf("User: " + string(testJson))

	document, exist := db.Login(t.Email, t.Password)

	if exist == false {
		http.Error(w, "User not exist", http.StatusBadRequest)
		log.Printf("User not exist")
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Error Generating JWT"+err.Error(), http.StatusBadRequest)
		return
	}

	response := model.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusCreated)

	expirationTime := time.Now().Add(time.Hour * 24)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
