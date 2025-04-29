package routers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"twitta/awsgo"
	"twitta/database"
	"twitta/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetImage(ctx context.Context, uploadType string, request events.APIGatewayProxyRequest, claim models.Claim) models.ApiResponse {

	var response models.ApiResponse           // create a new response object
	response.Status = 400                     // set the default status to 400
	ID := request.QueryStringParameters["id"] // get the ID from the query string parameters
	if len(ID) < 1 {                          // if the ID is empty, set the status to 400 and return the response
		response.Message = "El parámetro ID es obligatorio" // set the message to "The ID parameter is required"
		return response                                     // return the response
	}

	profile, err := database.SearchProfile(ID) // get the profile from the database using the ID
	if err != nil {                            // if there is an error in getting the profile, set the status to 500 and return the response
		response.Message = "Usuario no encontrado" + err.Error() // set the message to "User not found"
		return response                                          // return the response
	}
	var fileName string // create a variable to store the name of the file
	switch uploadType { // check the upload type
	case "avatar": // if the upload type is avatar, set the file name to the user ID and the file extension to .jpg
		fileName = profile.Avatar // set the file name to the user ID and the file extension to .jpg
	case "banner": // if the upload type is banner, set the file name to the user ID and the file extension to .jpg
		fileName = profile.Banner // set the file name to the user ID and the file extension to .jpg
	}

	fmt.Println("FileName", fileName) // print the file name to the console
	s3Conection := s3.NewFromConfig(awsgo.Cfg)

	file, err := downloadFromS3(ctx, s3Conection, fileName) // download the file from S3 using the file name

	if err != nil { // if there is an error in downloading the file, set the status to 500 and return the response
		response.Status = 500                                              // set the status to 500 (Internal Server Error)
		response.Message = "Error descargando archivo de S3" + err.Error() // set the message to "Error downloading file from S3"
		return response                                                    // return the response

	}

	response.CustomResp = &events.APIGatewayProxyResponse{ // create a new APIGatewayProxyResponse object
		StatusCode: 200,           // set the status code to 200 (OK)
		Body:       file.String(), // set the body to the file
		Headers: map[string]string{ // set the headers of the response
			"content-type":        "aplication/octet-stream",                            // set the content type to application/octet-stream
			"Content-Disposition": fmt.Sprintf("attachment; filename=\"%s\"", fileName), // set the content disposition to attachment and the file name to the file name

		},
	}

	return response

}

func downloadFromS3(ctx context.Context, s3Conection *s3.Client, fileName string) (*bytes.Buffer, error) { // downloadFromS3 is a function that downloads a file from S3 and returns the file as a byte buffer
	bucket := ctx.Value(models.Key("bucketName")).(string) // get the bucket name from the context
	object, err := s3Conection.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
	})

	if err != nil {
		return nil, err
	}

	defer object.Body.Close()
	fmt.Println("bucketName = " + bucket)

	file, err := io.ReadAll(object.Body)

	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(file)

	return buffer, nil

}
