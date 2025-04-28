package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweet(ID string, userID string) error {

	ctx := context.TODO()                          // Create a new context
	database := MongoClient.Database(DatabaseName) // Get the database
	collection := database.Collection("tweets")    // Get the collection

	tweetID, err := primitive.ObjectIDFromHex(ID) // Convert the ID string to an ObjectID
	if err != nil {
		return err // If there is an error, return the error
	}

	condition := bson.M{ // Create a new condition to find the tweet by ID and user ID
		"_id":    tweetID, // The ID of the tweet to be deleted
		"userID": userID,  // The ID of the user who created the tweet
	}

	_, err = collection.DeleteOne(ctx, condition) // Delete the tweet from the collection

	if err != nil {
		return err // If there is an error, return the error
	}

	return err // Return nil if there is no error
}
