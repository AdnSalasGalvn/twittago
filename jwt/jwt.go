package jwt

import (
	"context"
	"time"
	"twitta/models"

	"github.com/golang-jwt/jwt/v5"
)

func GennerateToken(ctx context.Context, user models.User) (string, error) { // This function generates a token for the user

	jwtSign := ctx.Value(models.Key("jwtsign")).(string) // get the jwt sign from the context
	myKey := []byte(jwtSign)                             // convert the jwt sign to a byte array

	payload := jwt.MapClaims{ // create a new payload for the token
		"email":     user.Email,                            // set the email of the user in the payload
		"name":      user.Name,                             // set the name of the user in the payload
		"lastName":  user.Lastnames,                        // set the last name of the user in the payload
		"birthDate": user.Birthday,                         // set the birth date of the user in the payload
		"biography": user.Biography,                        // set the biography of the user in the payload
		"location":  user.Location,                         // set the location of the user in the payload
		"website":   user.Website,                          // set the website of the user in the payload
		"_id":       user.ID.Hex(),                         // set the id of the user in the payload
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // set the expiration time of the token to 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) // create a new token with the payload and the signing method
	tokenString, err := token.SignedString(myKey)               // sign the token with the jwt sign and return the token string and the error
	if err != nil {                                             // if there is an error signing the token, return the error
		return tokenString, err // return the token string and the error
	}

	return tokenString, nil // return the token string and nil error
}
