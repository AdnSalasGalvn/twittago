package routers

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"mime"
	"mime/multipart"
	"strings"
	"twitta/database"
	"twitta/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type ReadSeeker struct { // ReadSeeker is a struct that implements the io.ReadSeeker interface
	io.Reader // Reader is an interface that allows reading from a stream
}

func (readSeekerModel *ReadSeeker) Seek(offset int64, whence int) (int64, error) { // Seek is a method that allows seeking to a specific position in the stream
	return 0, nil // return 0 and nil to indicate that the seek was successful

}

func UploadImage(ctx context.Context, uploadType string, request events.APIGatewayProxyRequest, claim models.Claim) models.ApiResponse { // UploadImage is a function that uploads an image to S3 and returns the URL of the image

	var response models.ApiResponse // create a new response object
	response.Status = 400           // set the default status to 400
	IDuser := claim.ID.Hex()        // get the ID of the user from the claim

	var fileName string  // create a variable to store the name of the file
	var user models.User // create a new user object

	bucket := aws.String(ctx.Value(models.Key("bucketName")).(string))

	switch uploadType { // check the upload type
	case "avatar": // if the upload type is avatar, set the file name to the user ID and the file extension to .png
		fileName = "avatar/" + IDuser + ".jpg" // set the file name to the user ID and the file extension to .png
		user.Avatar = fileName                 // set the avatar of the user to the file name
	case "banner": // if the upload type is banner, set the file name to the user ID and the file extension to .png
		fileName = "banners/" + IDuser + ".jpg" // set the file name to the user ID and the file extension to .png
		user.Banner = fileName                  // set the banner of the user to the file name
	}

	mediaType, params, err := mime.ParseMediaType(request.Headers["Content-Type"]) // parse the media type of the request
	if err != nil {                                                                // if there is an error in parsing the media type
		response.Status = 500          // set the status to 500 (Internal Server Error)
		response.Message = err.Error() // set the message to the error message
		return response                // return the response
	}

	if strings.HasPrefix(mediaType, "multipart/") { // if the media type is multipart

		body, err := base64.StdEncoding.DecodeString(request.Body) // decode the body of the request from base64
		if err != nil {                                            // if there is an error in decoding the body
			response.Status = 500
			response.Message = err.Error()
			return response
		}

		multipartReader := multipart.NewReader(bytes.NewReader(body), params["boundary"]) // create a new multipart reader with the body and the boundary
		part, err := multipartReader.NextPart()                                           // get the next part of the multipart reader

		if err != nil && err != io.EOF { // if there is an error in getting the next part
			response.Status = 500          // set the status to 500 (Internal Server Error)
			response.Message = err.Error() // set the status to 500 (Internal Server Error)
			return response
		}

		if err != io.EOF { // if there is no error in getting the next part
			if part.FileName() != "" { // if the file name is not empty
				buf := bytes.NewBuffer(nil)                   // create a new buffer to store the file
				if _, err := io.Copy(buf, part); err != nil { // copy the file to the buffer
					response.Message = err.Error() // set the message to the error message
					return response
				}

				amazonSession, err := session.NewSession(&aws.Config{ // create a new AWS session with the region and credentials
					Region: aws.String("us-east-1")}) // set the region to us-east-1

				if err != nil { // if there is an error in creating the session
					response.Status = 500          // set the status to 500 (Internal Server Error)
					response.Message = err.Error() // set the message to the error message
					return response
				}

				uploader := s3manager.NewUploader(amazonSession) // create a new S3 uploader with the session

				_, err = uploader.Upload(&s3manager.UploadInput{ // upload the file to S3
					Bucket: bucket,               // set the bucket to the bucket name
					Key:    aws.String(fileName), // set the key to the file name
					Body:   &ReadSeeker{buf},     // set the body to the buffer
				})

				if err != nil {
					response.Status = 500
					response.Message = err.Error()
					return response
				}
			}
		}

		status, err := database.ModifyRegister(user, IDuser)
		if err != nil || !status {
			response.Status = 400
			response.Message = "Error al modificar registro del usuario" + err.Error()
			return response
		}

	} else {
		response.Message = "Debe enviar una imagen con el 'Content-Type' de tipo 'multipart' en el Header"
		response.Status = 400
		return response
	}

	response.Status = 200
	response.Message = "Image Upload OK !"
	return response
}
