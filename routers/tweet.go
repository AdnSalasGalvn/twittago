package routers

import (
	"context"
	"encoding/json"
	"time"
	"twitta/database"
	"twitta/models"
)

func Tweet(ctx context.Context, claim models.Claim) models.ApiResponse { // This function handles the request and returns a response

	var tweetMessage models.Tweet // create a new tweet object

	var response models.ApiResponse // create a new response object

	response.Status = 400                              // set the default status to 400
	IDUser := claim.ID.Hex()                           // get the ID of the user from the claim
	body := ctx.Value(models.Key("body")).(string)     // get the body of the request from the context
	err := json.Unmarshal([]byte(body), &tweetMessage) // unmarshal the body into the tweet object
	if err != nil {                                    // if there is an error in unmarshaling the body
		response.Message = "Error al intentar decodificar el body" // set the message to an error message
		return response                                            // return the response
	}

	register := models.RecordTweet{ // create a new recordTweet object
		UserID:    IDUser,                                   // set the user ID to the ID of the user from the claim
		Message:   tweetMessage.Message,                     // set the message to the message from the tweet object
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"), // set the createdAt to the current date and time

	}

	_, status, err := database.InsertTweet(register) // insert the tweet into the database
	if err != nil {                                  // if there is an error in inserting the tweet
		response.Message = "Error al intentar insertar el tweet: " + err.Error() // set the message to an error message
		return response                                                          // return the response
	}

	if !status { // if the status is false, set the message to an error message
		response.Message = "Error al intentar insertar el registro" // set the message to an error message
		return response                                             // return the response
	}

	response.Status = 200                               // set the status to 200 (OK)
	response.Message = "Tweet registrado correctamente" // set the message to a success message
	return response                                     // return the response
}
