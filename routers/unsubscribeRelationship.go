package routers

import (
	"twitta/database"
	"twitta/models"

	"github.com/aws/aws-lambda-go/events"
)

func UnsubscribeRelationship(request events.APIGatewayProxyRequest, claim models.Claim) models.ApiResponse {
	var response models.ApiResponse
	response.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		response.Message = "El parámetro ID es obligatorio"
		return response
	}

	var relationship models.Relationship
	relationship.UserID = claim.ID.Hex()
	relationship.UserRelationshipID = ID

	status, err := database.EraseRelationship(relationship)
	if err != nil {
		response.Message = "Ocurrio un error al intentar borrar relación" + err.Error()
		return response
	}

	if !status {
		response.Message = "No se ha logrado borrar relación"
		return response
	}

	response.Status = 200
	response.Message = "Baja relacion OK!"
	return response
}
