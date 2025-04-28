package routers

import (
	"twitta/database"
	"twitta/models"

	"github.com/aws/aws-lambda-go/events"
)

func DeleteTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.ApiResponse {

	var response models.ApiResponse // create a new response object
	response.Status = 400           // set the default status to 400

	ID := request.QueryStringParameters["id"] // get the ID from the query string parameters

	if len(ID) < 1 {
		response.Message = "ID is required"
		return response
	}

	err := database.DeleteTweet(ID, claim.ID.Hex()) // call the deleteTweet function and get the error

	if err != nil { // call the deleteTweet function and check if it was successful
		response.Message = "Error al eliminar el tweet" + err.Error() // set the message to the error message
		return response
	}

	response.Status = 200                              // set the status to 200
	response.Message = "Tweet eliminado correctamente" // set the message to success message
	return response                                    // return the response

}
