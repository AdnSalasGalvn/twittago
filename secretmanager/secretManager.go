package secretmanager

import (
	"encoding/json"
	"fmt"
	"twitta/awsgo"
	"twitta/models"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws"
)

// Get the secret name and the secret data
func GetSecret(secretName string) (models.Secret, error) {
	var secretData models.Secret
	fmt.Println(" > pido secreto " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println(err.Error())
		return secretData, err

	}

	json.Unmarshal([]byte(*key.SecretString), &secretData)
	fmt.Println(" > lectura de Secret OK " + secretName)
	return secretData, nil

}
