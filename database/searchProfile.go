package database

import (
	"context"
	"twitta/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(ID string) (models.User, error) { // This function searches for a profile with the given ID and returns the profile and an error if there is any
	ctx := context.TODO()                          // create a new context
	database := MongoClient.Database(DatabaseName) // get the database from the MongoClient
	collection := database.Collection("users")     // get the collection from the database

	var profile models.User                        // create a new user model
	objectID, err := primitive.ObjectIDFromHex(ID) // convert the ID to an ObjectID

	if err != nil { // if there is an error converting the ID to an ObjectID, return the profile and the error
		return profile, err // return the profile and the error
	}

	condition := bson.M{
		"_id": objectID, // create a new condition for the query
	}

	err = collection.FindOne(ctx, condition).Decode(&profile) // find the profile with the given ID and decode it into the profile model

	profile.Password = "" // set the password to an empty string for security reasons

	if err != nil { // if there is an error finding the profile, return the profile and the error
		return profile, err // return the profile and the error
	}

	return profile, nil // return the profile and nil error
}
