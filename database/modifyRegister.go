package database

import (
	"context"
	"twitta/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyRegister(user models.User, ID string) (bool, error) {

	ctx := context.TODO()                      // Create a new context
	database := MongoClient.Database("twitta") // Get the database
	collection := database.Collection("users") // Get the collection

	register := make(map[string]interface{}) // Create a new map to store the user data
	if len(user.Name) > 0 {                  // If the name is not empty
		register["name"] = user.Name // Add the name to the map
	}

	if len(user.Lastnames) > 0 { // If the last name is not empty
		register["lastName"] = user.Lastnames // Add the last name to the map
	}

	register["birthday"] = user.Birthday // Add the birth date to the map

	if len(user.Avatar) > 0 { // If the avatar is not empty
		register["avatar"] = user.Avatar // Add the avatar to the map
	}

	if len(user.Banner) > 0 { // If the banner is not empty
		register["banner"] = user.Banner // Add the banner to the map
	}

	if len(user.Biography) > 0 { // If the biography is not empty
		register["biography"] = user.Biography // Add the biography to the map
	}

	if len(user.Location) > 0 { // If the location is not empty
		register["location"] = user.Location // Add the location to the map
	}

	if len(user.Website) > 0 { // If the website is not empty
		register["website"] = user.Website // Add the website to the map
	}

	updateString := bson.M{ // Create a new bson map to store the update data
		"$set": register, // Set the update data to the map
	}

	objectId, _ := primitive.ObjectIDFromHex(ID)     // Convert the ID string to an ObjectID
	filter := bson.M{"_id": bson.M{"$eq": objectId}} // Create a new filter to find the user by ID

	_, err := collection.UpdateOne(ctx, filter, updateString) // Update the user in the collection
	if err != nil {                                           // If there is an error in updating the user
		return false, err // Return false and the error
	}

	return true, nil // Return true and nil (no error)

}
