package main

import (
	"github.com/Elias2389/api-twitter/db"
	"github.com/Elias2389/api-twitter/handler"
	"log"
)

func main() {
	if db.ConnectionCheck() == 0 {
		log.Fatal("Don't connect")
	}

	handler.Handlers()
}
