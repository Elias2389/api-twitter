package handler

import (
	"github.com/Elias2389/api-twitter/midleware"
	"github.com/Elias2389/api-twitter/router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Handlers() {
	routers := mux.NewRouter()

	routers.HandleFunc("/register", midleware.CheckDb(router.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(routers)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
