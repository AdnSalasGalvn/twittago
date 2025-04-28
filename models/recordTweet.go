package models

type RecordTweet struct {
	UserID    string `bson:"userID" json:"userid,omitempty"`       // The ID of the user who created the tweet
	Message   string `bson:"message" json:"message,omitempty"`     // The message of the tweet
	CreatedAt string `bson:"createdAt" json:"createdAt,omitempty"` // The date and time when the tweet was created
}
