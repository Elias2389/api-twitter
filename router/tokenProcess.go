package router

import (
	"errors"
	"strings"

	"github.com/Elias2389/api-twitter/db"
	"github.com/Elias2389/api-twitter/model"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var UserId string

func tokenProcess(token string) (*model.Claim, bool, string, error) {
	myPass := []byte("key_of_pass")
	claims := &model.Claim{}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, "", errors.New("Token format invalid")
	}

	tk := strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myPass, nil
	})

	if err == nil {
		_, found, _ := db.CheckUserExist(claims.Email)

		if found == true {
			Email = claims.Email
			UserId = claims.ID.Hex()
		}
		return claims, found, UserId, nil
	}

	if !tkn.Valid {
		return claims, false, "", errors.New("Invalid token")
	}

	return claims, false, "", err
}
