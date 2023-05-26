package db

import (
	"fmt"
	"gambit/userLambda/models"
	"gambit/userLambda/tools"

	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Signing Up")
	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentence := "INSERT INTO users (User_UUID, User_Email, User_DateAdd) VALUES ('" + sig.UserUUID + "', '" + sig.UserEmail + "', '" + tools.DateMySQL() + "')"

	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return err
	}

	fmt.Println("User signed up")
	return nil
}
