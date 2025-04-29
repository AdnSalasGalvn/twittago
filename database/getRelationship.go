package database

import (
	"context"
	"twitta/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetRelationship(relationship models.Relationship) bool {
	ctx := context.TODO()
	database := MongoClient.Database(DatabaseName)
	collection := database.Collection("relationship")

	condition := bson.M{
		"userId":             relationship.UserID,
		"userRelationshipId": relationship.UserRelationshipID,
	}

	var results models.Relationship
	err := collection.FindOne(ctx, condition).Decode(&results)

	return err == nil
}
