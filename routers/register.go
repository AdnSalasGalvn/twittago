package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"twitta/database"
	"twitta/models"
)

func Register(ctx context.Context) models.ApiResponse {
	var userModel models.User
	var registerApiResponse models.ApiResponse
	registerApiResponse.Status = 400

	fmt.Println("Entre a Registro")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &userModel)

	if err != nil {
		registerApiResponse.Message = err.Error()
		fmt.Println(registerApiResponse.Message)
		return registerApiResponse
	}

	if len(userModel.Email) == 0 {
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

	_, userFounded, _ := database.CheckUserAlreadyExist(userModel.Email)
	if userFounded {
		registerApiResponse.Message = "Ya existe un usuario registrado con ese Email"
		fmt.Println(registerApiResponse.Message)
		return registerApiResponse

	}

	_, status, err := database.InsertRecord(userModel)
	if err != nil {
		registerApiResponse.Message = "Ocurrió un error al intentar realizar el registro de usuario" + err.Error()
		fmt.Println(registerApiResponse.Message)
		return registerApiResponse
	}

	if !status {
		registerApiResponse.Message = "No se ha logrado insertar el registro del usuario"
		fmt.Println(registerApiResponse.Message)
		return registerApiResponse
	}

	registerApiResponse.Status = 200
	registerApiResponse.Message = "Registro OK"

	fmt.Println(registerApiResponse.Message)
	return registerApiResponse
}
