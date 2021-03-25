package db

import "golang.org/x/crypto/bcrypt"

func EncryptPass(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(bytes), err
}
