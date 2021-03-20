package router

import (
	"encoding/json"
	"github.com/Elias2389/api-twitter/db"
	"github.com/Elias2389/api-twitter/model"
	"net/http"
)

func Register(writter http.ResponseWriter, request *http.Request) {
	var t model.User
	err := json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		http.Error(writter, "Error in data required" + err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(writter, "Error in email required", http.StatusBadRequest)
		return
	}

	if len(t.Password) == 0 {
		http.Error(writter, "Password must have more than 3 ", http.StatusBadRequest)
		return
	}

	_, finded, _ := db.CheckUserExist(t.Email)
	if finded == true {
		http.Error(writter, "User exist" + err.Error(), http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertRegister(t)
	if err != nil {
		http.Error(writter, "Error with insert register" + err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(writter, "Error with insert register Status ", http.StatusBadRequest)
		return
	}

	writter.WriteHeader(http.StatusCreated)
}