package midleware

import (
	"net/http"
	//"github.com/Elias2389/api-twitter/router"
)

func validateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//_, _, _, err := router
	}
}
