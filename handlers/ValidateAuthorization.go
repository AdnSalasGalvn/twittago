package handlers

import (
	"context"
	"fmt"
	"twitta/jwt"
	"twitta/models"

	"github.com/aws/aws-lambda-go/events"
)

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
