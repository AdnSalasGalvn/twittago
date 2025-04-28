package database

import (
	"context"
	"twitta/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTweets(ID string, page int64) ([]*models.RecordTweet, bool) {
	ctx := context.TODO()                          // Create a new context
	database := MongoClient.Database(DatabaseName) // Get the database
	collection := database.Collection("tweets")    // Get the collection

	var results []*models.RecordTweet // Create a new slice to store the results

	condition := bson.M{
		"userID": ID,
	} // Create a new condition to find the tweets by user ID

	options := options.Find()                              // Create a new options object to set the options for the query
	options.SetLimit(20)                                   // Set the limit to 20 tweets
	options.SetSort(bson.D{{Key: "createdAt", Value: -1}}) // Sort the tweets by createdAt in descending order
	options.SetSkip((page - 1) * 20)                       // Set the skip to the page number minus 1 times 20

	cursor, err := collection.Find(ctx, condition, options) // Find the tweets in the collection
	if err != nil {
		return results, false // If there is an error, return the results and false
	}

	for cursor.Next(ctx) { // Iterate over the cursor
		var record models.RecordTweet // Create a new RecordTweet object to store the tweet
		err := cursor.Decode(&record) // Decode the tweet into the RecordTweet object
		if err != nil {               // If there is an error in decoding the tweet
			return results, false // Return the results and false
		}
		results = append(results, &record) // Append the tweet to the results slice
	}
	return results, true // Return the results and true

}
