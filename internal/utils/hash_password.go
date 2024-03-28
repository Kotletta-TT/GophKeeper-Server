package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(hashedPassword, password string) error {
	byteHash := []byte(hashedPassword)
	bytePwd := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return err
	}
	return nil
}
