package db

import (
	"context"
	"github.com/Elias2389/api-twitter/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func InsertRegister(u model.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dataBase := MongoConnect.Database("twittor")
	col := dataBase.Collection("user")

	u.Password, _ = EncryptPass(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)
	return objId.String(), true, nil
}
