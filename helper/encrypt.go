package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func PasswordBecrypt(password string) (response string, err error) {
	password_bcrypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "error", err
	}
	return string(password_bcrypt), nil
}

func PasswordMatch(password string, password_byte_hash []byte) bool {
	byte_password := []byte(password)
	if err := bcrypt.CompareHashAndPassword(password_byte_hash, byte_password); err != nil {
		return false
	}
	return true
}
