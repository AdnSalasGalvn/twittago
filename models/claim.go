package models

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claim struct { // estructura que representa el payload del token JWT
	Email                string              `json:"email"`                    // email del usuario
	ID                   *primitive.ObjectID `bson:"_id" json:"_id,omitempty"` // id del usuario
	jwt.RegisteredClaims                     // claims registrados del token
}
