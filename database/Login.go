package database

import (
	"twitta/models"

	"golang.org/x/crypto/bcrypt"
)

func TryToLogin(email string, password string) (models.User, bool) { // This function tries to login the user with the email and password

	user, exist, _ := CheckUserAlreadyExist(email) // check if the user already exists in the database

	if !exist { // if the user does not exist, return an error and a message
		return user, false // return the error and the message
	}

	passwordBytes := []byte(password)   // convert the password to a byte array
	passwordDB := []byte(user.Password) // convert the password from the database to a byte array

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes) // compare the password from the database with the password from the request

	if err != nil { // if there is an error comparing the passwords, return an error and a message
		return user, false // return the error and the message
	} // return the user and true if the passwords match

	return user, true // return the user and true if the passwords match

}
