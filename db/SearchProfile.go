package db

import (
	"context"
	"github.com/Elias2389/api-twitter/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func FindProfileById(ID string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoConnect.Database("twittor")
	col := db.Collection("user")

	var profile model.User

	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		log.Print("Register not found" + err.Error())
		return profile, err
	}

	return profile, nil
}
