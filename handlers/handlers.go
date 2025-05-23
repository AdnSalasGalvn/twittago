package handlers

import (
	"context"
	"twitta/models"
	"twitta/routers"

	"github.com/aws/aws-lambda-go/events"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.ApiResponse { // This function handles the request and returns a response

	var response models.ApiResponse // create a new response object
	response.Status = 400           // set the default status to 400

	isOk, statusCode, msg, claim := ValidateAuthorization(ctx, request) // validate the authorization of the request

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
		case "login": // if the path is login, call the login function
			return routers.Login(ctx) // call the login function and return the response
		case "tweet": // if the path is tweet, call the tweet function
			return routers.Tweet(ctx, claim) // call the tweet function and return the response
		case "registerrelationship": // if the path is tweet, call the tweet function
			return routers.RegisterRelationship(ctx, request, claim) // call the tweet function and return the response
		case "uploadAvatar": // if the path is uploadAvatar, call the uploadAvatar function´
			return routers.UploadImage(ctx, "avatar", request, claim) // call the uploadAvatar function and return the response
		case "uploadBanner": // if the path is uploadBanner, call the uploadBanner function
			return routers.UploadImage(ctx, "banner", request, claim) // call the uploadBanner function and return the response
		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "seeprofile": // if the path is seeProfile, call the seeProfile function
			return routers.SeeProfile(request) // call the seeProfile function and return the response
		case "readtweets": // if the path is readTweet, call the readTweet function
			return routers.ReadTweets(request) // call the readTweet function and return the response
		case "getAvatar": // if the path is getAvatar, call the getAvatar function
			return routers.GetImage(ctx, "avatar", request, claim) // call the getAvatar function and return the response
		case "getBanner": // if the path is getBanner, call the getBanner function
			return routers.GetImage(ctx, "banner", request, claim) // call the getBanner function and return the response
		case "getrelationship": //
			return routers.GetRelationship(request, claim)
		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		case "modifyprofile": // if the path is modifyProfile, call the modifyProfile function
			return routers.ModifyProfile(ctx, claim) // call the modifyProfile function and return the response

		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		case "deletetweet": // if the path is deleteTweet, call the deleteTweet function
			return routers.DeleteTweet(request, claim) // call the deleteTweet function and return the response
		case "unsubscriberelationship": // if the path is deleteTweet, call the deleteTweet function
			return routers.UnsubscribeRelationship(request, claim) // call the deleteTweet function and return the response
		}
	}

	// if the method is not valid, return an error and a message
	response.Message = "Method Invalid"
	return response
}
