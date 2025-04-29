package routers

import (
	"encoding/json"
	"twitta/database"
	"twitta/models"

	"github.com/aws/aws-lambda-go/events"
)

func GetRelationship(request events.APIGatewayProxyRequest, claim models.Claim) models.ApiResponse {

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

	var responseQueryRelation models.ResponseQueryRelation

	existRelationship := database.GetRelationship(relationship)

	if !existRelationship {
		responseQueryRelation.Status = false
	} else {
		responseQueryRelation.Status = true
	}

	jsonResponse, err := json.Marshal(existRelationship)
	if err != nil {
		response.Status = 500
		response.Message = "Error al formatear los datos de los usuarios como JSON " + err.Error()
		return response
	}

	response.Status = 200
	response.Message = string(jsonResponse)

	return response

}
