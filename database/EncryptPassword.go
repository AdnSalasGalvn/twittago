package database

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	if err != nil {
		return err.Error(), err
	}
	return string(bytes), nil
}
