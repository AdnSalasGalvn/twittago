package jwt

import (
	"errors"
	"fmt"
	"strings"
	"twitta/models"

	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUser string

func ProcessToken(tk string, JWTSign string) (*models.Claim, bool, string, error) { // función que procesa el token JWT

	myKey := []byte(JWTSign) // llave secreta para firmar el token

	var claims models.Claim // claims es el payload del token

	splitToken := strings.Split(tk, "Bearer") // separa el token del tipo de token (Bearer)

	fmt.Println("Token Dividido:", splitToken, " Longitud:", len(splitToken[1])) //	Valor y longitud del token

	if len(splitToken) != 2 { // si no hay un token o el formato es incorrecto
		return &claims, false, string(""), errors.New("formato de token inválido") // retorna error
	}

	tk = strings.TrimSpace(splitToken[1]) // quita los espacios en blanco del token

	fmt.Println(
		"(Archivo ProcessToken.go) Valor de Token:", tk, "longitud:",
		len(splitToken[1]),
	) // imprime el token y su longitud

	token, err := jwt.ParseWithClaims(
		tk, &claims,
		func(token *jwt.Token) (any, error) { // parsea el token y lo valida
			return myKey, nil // retorna la llave secreta
		},
	)
	if err != nil {
		fmt.Printf(
			"Ocurrio un error el intentar parsear el token: %s \n", err.Error(),
		)
		return &claims, false, string(""), err
	}

	if !token.Valid { // si el token no es válido
		return &claims, false, string(""), errors.New("token Inválido") // retorna error
	}

	return &claims, true, string(""), err // retorna el token y el error
}
