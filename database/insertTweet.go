package database

import (
	"context"
	"twitta/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(tweet models.RecordTweet) (string, bool, error) { // This function inserts a tweet into the database and returns the ID of the tweet, a status, and an error if any
	ctx := context.TODO()                          // create a new context
	database := MongoClient.Database(DatabaseName) // get the database from the MongoClient
	colletion := database.Collection("tweets")     // get the collection from the database

	register := bson.M{ // create a new bson map to store the tweet data
		"userID":    tweet.UserID,    // set the user ID to the ID of the user from the tweet object
		"message":   tweet.Message,   // set the message to the message from the tweet object
		"createdAt": tweet.CreatedAt, // set the createdAt to the current date and time
	}
	result, err := colletion.InsertOne(ctx, register) // insert the tweet into the collection
	if err != nil {                                   // if there is an error in inserting the tweet
		return "", false, err // return an empty string, false, and the error
	}

	objectID, _ := result.InsertedID.(primitive.ObjectID) // get the ID of the inserted tweet
	return objectID.String(), true, nil                   // return the ID of the tweet as a string, true, and nil (no error)

}
