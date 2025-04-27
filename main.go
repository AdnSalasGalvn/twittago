package main

import (
	"context"
	"os"
	"strings"
	"twitta/awsgo"
	"twitta/database"
	"twitta/handlers"
	"twitta/models"
	"twitta/secretmanager"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {

	lambda.Start(LambdaExecution)
}

func LambdaExecution(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) { // This function is the entry point for the Lambda function. It receives the request and returns a response.
	var res *events.APIGatewayProxyResponse // create a new response object

	awsgo.InitializeAWS() // initialize the AWS SDK

	if !ValidateParameters() { // check if the environment variables are set
		res = &events.APIGatewayProxyResponse{ // if not, return an error and a message
			StatusCode: 400,                                                                                        // set the status code to 400
			Body:       "Error en las variables de entorno. Deben incluir 'ScretName', 'BucketName', 'UrlPrefix' ", // set the body to the error message
			Headers: map[string]string{ // set the headers to indicate that the response is in JSON format
				"Content-Type": "application/json", // set the content type to JSON
			},
		}
		return res, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName")) // get the secret from AWS Secrets Manager using the secret name from the environment variable

	if err != nil { // if there is an error getting the secret, return an error and a message
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura de Secret: " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["twittago"], os.Getenv("UrlPrefix"), "", -1) // Remove the prefix from the path
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)                          // Set the path in the context
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)          // Set the method in the context
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)          // Set the user in the context
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)      // Set the password in the context
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)              // Set the host in the context
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)      // Set the database in the context
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtsign"), SecretModel.JWTSign)        // Set the JWT sign in the context
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)                  // Set the body in the context
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName")) // Set the bucket name in the context

	//checking Database connection and connected with Database.

	err = database.ConnectDB(awsgo.Ctx) // connect to the database using the credentials from the context

	if err != nil { // if there is an error connecting to the database, return an error and a message
		res = &events.APIGatewayProxyResponse{ // set the status code to 500
			StatusCode: 500,                                      // set the status code to 500
			Body:       "Error conectando la DB: " + err.Error(), // set the body to the error message
			Headers: map[string]string{ // set the headers to indicate that the response is in JSON format
				"Content-Type": "application/json", // set the content type to JSON
			},
		}
		return res, nil
	}

	APIResponse := handlers.Handlers(awsgo.Ctx, request) // call the Handlers function to process the request and get the response

	if APIResponse.CustomResp == nil { // if the response is nil, return an error and a message
		res = &events.APIGatewayProxyResponse{
			StatusCode: APIResponse.Status,
			Body:       APIResponse.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	} else { // if the response is not nil, return the response
		return APIResponse.CustomResp, nil
	}
}

func ValidateParameters() bool { // This function checks if the environment variables are set
	_, bringParameters := os.LookupEnv("SecretName")
	if !bringParameters {
		return bringParameters
	}

	_, bringParameters = os.LookupEnv("BucketName")
	if !bringParameters {
		return bringParameters
	}

	_, bringParameters = os.LookupEnv("UrlPrefix")
	if !bringParameters {
		return bringParameters
	}

	return bringParameters // return true if all the environment variables are set

}
