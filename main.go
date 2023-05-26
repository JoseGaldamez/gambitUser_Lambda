package main

import (
	"context"
	"errors"
	"fmt"
	"gambit/userLambda/awsgo"
	"gambit/userLambda/db"
	"gambit/userLambda/models"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(runUserLambda)
}

func runUserLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	fmt.Println("-> Init UserLambda")

	awsgo.InitializeAWS()

	fmt.Println("-> Initialize AWS - OK")

	if !validateParamenters() {
		fmt.Println("Missing environment parameter: 'SecretName'")
		error := errors.New("missing environment parameters: 'SecretName'")
		return event, error
	}
	fmt.Println("-> ValidateParameters - OK")

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("UserEmail: " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Sub: " + data.UserUUID)
		}
	}
	fmt.Println("-> Atributes event.Request - OK")

	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error reading secret: " + err.Error())
		return event, err
	}

	fmt.Println("-> ReadSecret - OK")

	err = db.SignUp(data)
	if err != nil {
		fmt.Println("Error signing up: " + err.Error())
		return event, err
	}

	fmt.Println("-> db.SignUp - OK")

	return event, nil
}

func validateParamenters() bool {
	_, existSecrectName := os.LookupEnv("SecretName")
	return existSecrectName
}
