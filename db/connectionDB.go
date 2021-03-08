package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const userMongoDb string = ""
const passMongoDb string = ""

var MongoConnect = connectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://" + userMongoDb + ":" + passMongoDb + "@twittor.blv2u.mongodb.net/test")

func connectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Printf("Connected to DB")
	return client
}

func ConnectionCheck() int {
	err := MongoConnect.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
