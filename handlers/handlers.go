package handlers

import (
	"context"
	"fmt"
	"twitta/jwt"
	"twitta/models"
	"twitta/routers"

	"github.com/aws/aws-lambda-go/events"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.ApiResponse { // This function handles the request and returns a response

	var response models.ApiResponse // create a new response object
	response.Status = 400           // set the default status to 400

	isOk, statusCode, msg, _ := ValidateAuthorization(ctx, request) // validate the authorization of the request

	if !isOk { // if the authorization is not valid, return an error and a message
		response.Status = statusCode // set the status code to the one returned by the validation function
		response.Message = msg       // set the message to the one returned by the validation function
		return response
	}

	switch ctx.Value(models.Key("method")).(string) { // check the method of the request

	case "POST": // if the method is POST, check the path of the request
		switch ctx.Value(models.Key("path")).(string) { // check the path of the request
		case "register": // if the path is register, call the register function
			return routers.Register(ctx) // call the register function and return the response
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
	// if the method is not valid, return an error and a message
	response.Message = "Method Invalid"
	return response
}

func ValidateAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) { // This function validates the authorization of the request
	// check if the path is valid and if the method is valid

	path := ctx.Value(models.Key("path")).(string)
	if path == "register" || path == "login" || path == "getAvatar" || path == "getBanner" { // if the path is valid, return true
		return true, 200, "", models.Claim{} // return true
	}

	token := request.Headers["Authorization"] // get the token from the request headers

	if len(token) == 0 { // if the token is empty, return false
		return false, 401, "Token requerido", models.Claim{} // return false
	}

	claim, allOK, msg, err := jwt.ProcessToken(token, ctx.Value(models.Key("jwtsign")).(string)) // process the token and get the claims

	if !allOK { // if the token is not valid, return false
		if err != nil {
			fmt.Println("Error en el token: " + err.Error()) // print the error
			fmt.Println("NO AllOk: err != nil")
			return false, 401, err.Error(), models.Claim{} // return false
		} else { // if err is nil, return false
			fmt.Println("Error en el token: " + msg) // print the error
			fmt.Println("NO AllOk: Else")
			return false, 401, msg, models.Claim{} // return false, msg is the error message, claim is the claims, which is empty
		}
	}

	// end of the token validation, token is valid and we can use the claim
	fmt.Println("Token Ok")
	return true, 200, msg, *claim
}
