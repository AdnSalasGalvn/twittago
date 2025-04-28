package routers

import (
	"encoding/json"
	"fmt"
	"twitta/database"
	"twitta/models"

	"github.com/aws/aws-lambda-go/events"
)

func SeeProfile(request events.APIGatewayProxyRequest) models.ApiResponse { // This function handles the seeProfile request and returns a response

	var response models.ApiResponse // create a new response object
	response.Status = 400           // set the default status to 400

	fmt.Println("Entré en seeProfile")        // print a message to the console
	ID := request.QueryStringParameters["id"] // get the id from the query string parameters
	if len(ID) < 1 {                          // check if the id is empty
		response.Message = "ID del usuario es obligatorio" // set the message to the one returned by the validation function"
		return response                                    // return the error and the message}
	}

	profile, err := database.SearchProfile(ID) // search for the profile with the id
	if err != nil {                            // if there is an error searching for the profile, return an error and a message
		response.Message = "Error al buscar el perfil > " + err.Error() // set the message to the one returned by the validation function"
		return response                                                 // return the error and the message
	}

	responseJson, err := json.Marshal(profile) // marshal the profile into a byte array
	if err != nil {                            // if there is an error marshaling the profile, return an error and a message
		response.Status = 500                                                                 // set the status code to 500
		response.Message = "Error al formatear los datos del usuario como JSON" + err.Error() // set the message to the one returned by the validation function"
	}

	response.Status = 200                   // set the status code to 200
	response.Message = string(responseJson) // set the message to the profile in JSON format
	return response                         // return the response
}
