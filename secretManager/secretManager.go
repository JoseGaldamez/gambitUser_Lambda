package secretmanager

import (
	"encoding/json"
	"fmt"
	"gambit/userLambda/awsgo"
	"gambit/userLambda/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var secretRDSJson models.SecretRDSJson

	fmt.Println("> Ask for secret: " + secretName)
	secretManageClient := secretsmanager.NewFromConfig(awsgo.Cfg)

	key, err := secretManageClient.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{SecretId: aws.String(secretName)})

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return secretRDSJson, err
	}

	json.Unmarshal([]byte(*key.SecretString), &secretRDSJson)

	return secretRDSJson, nil
}
