package database

import (
	"context"
	"twitta/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRecord(u models.User) (string, bool, error) {
	ctx := context.TODO()
	database := MongoClient.Database(DatabaseName)

	databasecollection := database.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)
	result, err := databasecollection.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
