package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"twitta/database"
	"twitta/models"
)

func Register(ctx context.Context) models.ApiResponse { // This function handles the registration of a new user
	var userModel models.User                  // create a new user model
	var registerApiResponse models.ApiResponse // create a new response object
	registerApiResponse.Status = 400           // set the default status to 400

	fmt.Println("Entre a Registro")

	body := ctx.Value(models.Key("body")).(string)  // get the body of the request from the context
	err := json.Unmarshal([]byte(body), &userModel) // unmarshal the body into the user model

	if err != nil { // if there is an error unmarshaling the body, return an error and a message
		registerApiResponse.Message = err.Error()
		fmt.Println(registerApiResponse.Message)
		return registerApiResponse
	}

	if len(userModel.Email) == 0 { // check if the email is empty
		registerApiResponse.Message = "Debes especificar el Email"
		fmt.Println(registerApiResponse.Message)
		return registerApiResponse
	}

	// check that the password is longer than 6 characters.
	if len(userModel.Password) < 6 {
		registerApiResponse.Message = "Debe especificar una contraseña de al menos 6 carácteres"
		fmt.Println(registerApiResponse.Message)
		return registerApiResponse
	}

	_, userFounded, _ := database.CheckUserAlreadyExist(userModel.Email) // check if the user already exists in the database
	if userFounded {                                                     // if the user already exists, return an error and a message
		registerApiResponse.Message = "Ya existe un usuario registrado con ese Email"
		fmt.Println(registerApiResponse.Message)
		return registerApiResponse

	}

	_, status, err := database.InsertRecord(userModel) // insert the user into the database
	if err != nil {                                    // if there is an error inserting the user, return an error and a message
		registerApiResponse.Message = "Ocurrió un error al intentar realizar el registro de usuario" + err.Error() // return an error and a message
		fmt.Println(registerApiResponse.Message)                                                                   // print the error message
		return registerApiResponse                                                                                 // return the error and the message
	}

	if !status { // if the status is false, return an error and a message
		registerApiResponse.Message = "No se ha logrado insertar el registro del usuario" // return an error and a message
		fmt.Println(registerApiResponse.Message)                                          // print the error message
		return registerApiResponse                                                        // return the error and the message
	}

	registerApiResponse.Status = 200            // set the status to 200
	registerApiResponse.Message = "Registro OK" // set the message to "OK"

	fmt.Println(registerApiResponse.Message) // print the message
	return registerApiResponse               // return the response object
}
