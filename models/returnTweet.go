package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReturnTweet struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`              // The ID of the tweet
	UserID    string             `bson:"userid" json:"userid,omitempty"`       // The ID of the user who created the tweet
	Message   string             `bson:"message" json:"message,omitempty"`     // The message of the tweet
	CreatedAt string             `bson:"createdAt" json:"createdAt,omitempty"` // The date and time when the tweet was created
}
