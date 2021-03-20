package db

import (
	"context"
	"github.com/Elias2389/api-twitter/model"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func CheckUserExist(email string) (model.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoConnect.Database("twittor")
	col := db.Collection("user")

	condition := bson.M{ "email": email }


	var result model.User
	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}


