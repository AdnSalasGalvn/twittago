package models

type Tweet struct {
	Message string `bson:"message" json:"message"` // The message of the tweet
}
