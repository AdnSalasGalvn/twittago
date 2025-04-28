package routers

import (
	"encoding/json"
	"strconv"
	"twitta/database"
	"twitta/models"

	"github.com/aws/aws-lambda-go/events"
)

func ReadTweets(request events.APIGatewayProxyRequest) models.ApiResponse {
	var response models.ApiResponse // create a new response object
	response.Status = 400           // set the default status to 400

	ID := request.QueryStringParameters["id"]     // get the ID from the query string parameters
	page := request.QueryStringParameters["page"] // get the page from the query string parameters

	if len(ID) < 1 {
		response.Message = "ID is required"
		return response
	}

	if len(page) < 1 {
		page = "1"
	}

	pageIntegerValue, err := strconv.Atoi(page) // convert the page string to an int

	if err != nil {
		response.Message = "Debe enviar un parametro page como un valor mayor a 0"
		return response
	}

	tweets, ok := database.ReadTweets(ID, int64(pageIntegerValue)) // call the readTweet function and get the tweets

	if !ok { // if the readTweet function returns false, set the message to the error message and return the response
		response.Message = "Error al leer los tweets"
		return response
	}

	responseJson, err := json.Marshal(tweets) // marshal the tweets to JSON
	if err != nil {                           // if there is an error in marshaling the tweets, set the message to the error message and return the response
		response.Status = 500 // set the status to 500
		response.Message = "Error al formaterar los datos de los usuarios como JSON"
		return response
	}

	response.Status = 200                   // set the status to 200
	response.Message = string(responseJson) // set the message to the JSON string of the tweets
	return response                         // return the response

}
