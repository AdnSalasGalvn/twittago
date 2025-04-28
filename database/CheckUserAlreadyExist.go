package database

import (
	"context"
	"twitta/models"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckUserAlreadyExist(email string) (models.User, bool, string) {

	ctx := context.TODO()                          // context for the database connection
	database := MongoClient.Database(DatabaseName) // database connection

	databasecollection := database.Collection("users") // collection name
	condition := bson.M{"email": email}                // condition to find the user
	var result models.User                             // create a new user model

	err := databasecollection.FindOne(ctx, condition).Decode(&result) // find the user by email and decode the result into the user model
	ID := result.ID.Hex()                                             // get the ID of the user
	if err != nil {                                                   // if there is an error finding the user,
		return result, false, ID // return the user model, false and the ID
	}
	// if the user exists, return the user model, true and the ID
	return result, true, ID
}
