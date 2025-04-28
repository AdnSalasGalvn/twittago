package routers

import (
	"context"
	"encoding/json"
	"twitta/database"
	"twitta/models"
)

func ModifyProfile(ctx context.Context, claim models.Claim) models.ApiResponse { // ModifyProfile modifies the profile of the user

	var response models.ApiResponse // Create a new ApiResponse object
	response.Status = 400           // Set the default status to 400 (Bad Request)

	var user models.User // Create a new User object

	body := ctx.Value(models.Key("body")).(string) // Get the request body from the context

	err := json.Unmarshal([]byte(body), &user) // Unmarshal the request body into the User object
	if err != nil {                            // If there is an error in unmarshalling
		response.Message = "Datos incorrectos" // Set the error message
	}

	status, err := database.ModifyRegister(user, claim.ID.Hex()) // Call the ModifyRegister function to update the user profile
	if err != nil {                                              // If there is an error in updating the profile
		response.Message = "Ocurrió un error al intentar modificar el registro" // Set the error message
		return response
	}

	if !status { // If the status is false
		response.Message = "No se pudo modificar el registro del usuario" // Set the error message
		return response
	}
	response.Status = 200                                  // Set the status to 200 (OK)
	response.Message = "Registro modificado correctamente" // Set the success message

	return response // Return the response
}
