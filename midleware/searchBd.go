package midleware

import (
	"github.com/Elias2389/api-twitter/db"
	"net/http"
)

func CheckDb(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if db.ConnectionCheck() == 0 {
			http.Error(writer, "Lost connection with DB", 500)
		}
		next.ServeHTTP(writer, request)
	}
}
