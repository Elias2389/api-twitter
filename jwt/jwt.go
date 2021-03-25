package jwt

import (
	"github.com/Elias2389/api-twitter/model"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJWT(t model.User) (string, error) {
	myPass := []byte("key_of_pass")
	payload := jwt.MapClaims{
		"email": t.Email,
		"name":  t.Name,
		"date":  t.Date,
		"_id":   t.ID.Hex(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(myPass)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
