package database

import (
	"context"
	"twitta/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRecord(userModel models.User) (string, bool, error) {
	ctx := context.TODO()                          // context for the database connection
	database := MongoClient.Database(DatabaseName) // database connection

	databasecollection := database.Collection("users") // collection name

	// create a new collection for the user
	var err error                                                 // variable to store the error
	userModel.Password, err = EncryptPassword(userModel.Password) // encrypt the password

	if err != nil { // if there is an error encrypting the password, return an error and a message
		return "", false, err
	}

	result, err := databasecollection.InsertOne(ctx, userModel)

	if err != nil { // if there is an error inserting the user, return an error and a message
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID) // get the ID of the inserted user
	return ObjID.String(), true, nil
}
