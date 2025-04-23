package handlers

import (
	"context"
	"fmt"
	"twitta/jwt"
	"twitta/models"
	"twitta/routers"

	"github.com/aws/aws-lambda-go/events"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.ApiResponse {
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var response models.ApiResponse
	response.Status = 400

	isOk, statusCode, msg, _ := ValidateAuthorization(ctx, request)

	if !isOk {
		response.Status = statusCode
		response.Message = msg
		return response
	}

	switch ctx.Value(models.Key("method")).(string) {

	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "register":
			return routers.Register(ctx)
		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
	}
	response.Message = "Method Invalid"
	return response
}

func ValidateAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) { // 	función que valida la autorización del token JWT

	path := ctx.Value(models.Key("path")).(string)
	if path == "register" || path == "login" || path == "getAvatar" || path == "getBanner" { // si es un path que no requiere autorización
		return true, 200, "", models.Claim{} // retorna true
	}

	token := request.Headers["Authorization"]

	fmt.Println("(Archivo handlers.go) Valor del token recibido:", token)

	if len(token) == 0 { // si no hay token
		return false, 401, "Token requerido", models.Claim{} // retorna error
	}

	claim, allOK, msg, err := jwt.ProcessToken(token, ctx.Value(models.Key("jwtsign")).(string)) // llama a la función ProcessToken

	if !allOK { // si no es válido el token
		if err != nil { // si hay un error en el token
			fmt.Println("Error en el token: " + err.Error()) // imprime el error
			fmt.Println("NO AllOk: err != nil")              // imprime el error
			return false, 401, err.Error(), models.Claim{}   // retorna error
		} else { // si el token no es válido
			fmt.Println("Error en el token: " + msg) // 	imprime el error
			fmt.Println("NO AllOk: Else")            // imprime el error
			return false, 401, msg, models.Claim{}   // retorna error
		}
	}

	fmt.Println("Token Ok")       // imprime el token
	return true, 200, msg, *claim // retorna el token y el error
}
