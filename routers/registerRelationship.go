package routers

import (
	"context"
	"twitta/database"
	"twitta/models"

	"github.com/aws/aws-lambda-go/events"
)

func RegisterRelationship(ctx context.Context, request events.APIGatewayProxyRequest, claim models.Claim) models.ApiResponse {
	var response models.ApiResponse
	response.Status = 400

	ID := request.QueryStringParameters["id"]

	if len(ID) < 1 {
		response.Message = "El primer parámetro ID es Obligatorio"
		return response
	}

	var relationship models.Relationship
	relationship.UserID = claim.ID.Hex()
	relationship.UserRelationshipID = ID

	status, err := database.InsertRelationship(relationship)

	if err != nil {
		response.Message = "Ocurrió un error al intentar insertar relación " + err.Error()
		return response
	}

	if !status {
		response.Message = "No se ha logrado insertar la realación "
		return response
	}

	response.Status = 200
	response.Message = "Alta de relación OK"
	return response

}
