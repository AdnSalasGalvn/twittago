package database

import (
	"context"
	"twitta/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRecord(u models.User) (string, bool, error) {
	ctx := context.TODO()                          // context for the database connection
	database := MongoClient.Database(DatabaseName) // database connection

	databasecollection := database.Collection("users") // collection name

	// create a new collection for the user
	u.Password, _ = EncryptPassword(u.Password) // encrypt the password
	result, err := databasecollection.InsertOne(ctx, u)

	if err != nil { // if there is an error inserting the user, return an error and a message
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID) // get the ID of the inserted user
	return ObjID.String(), true, nil
}
