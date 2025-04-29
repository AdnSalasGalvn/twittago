package database

import (
	"context"
	"twitta/models"
)

func InsertRelationship(relationship models.Relationship) (bool, error) {
	ctx := context.TODO()

	database := MongoClient.Database(DatabaseName)
	collection := database.Collection("relationship")

	_, err := collection.InsertOne(ctx, relationship)
	if err != nil {
		return false, err
	}

	return true, nil
}
