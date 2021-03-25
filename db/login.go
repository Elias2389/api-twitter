package db

import (
	"github.com/Elias2389/api-twitter/model"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Login(email string, pass string) (model.User, bool) {
	user, found, _ := CheckUserExist(email)

	if found == false {
		return user, false
	}

	passwordBytes := []byte(pass)
	passDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passDB, passwordBytes)
	if err != nil {
		log.Fatal(err)
		return user, false
	}

	return user, true
}
