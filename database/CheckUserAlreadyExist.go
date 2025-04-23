package database

import (
	"context"
	"twitta/models"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckUserAlreadyExist(email string) (models.User, bool, string) {

	ctx := context.TODO()
	database := MongoClient.Database(DatabaseName)

	databasecollection := database.Collection("users")
	condition := bson.M{"email": email}
	var result models.User

	err := databasecollection.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
