package models

type Relationship struct {
	UserID             string `bson:"userId" json:"userId"`
	UserRelationshipID string `bson:"userRelationshipId" json:"userRelationshipId"`
}
