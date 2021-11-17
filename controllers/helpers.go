package controllers

import (
	"errors"
	"log"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func IsPasswordStrong(password string) (bool, error) {
	var IsLength, IsUpper, IsLower, IsNumber, IsSpecial bool
	if len(password) < 8 {
		return false, errors.New("password length should be more than 8")
	}
	IsLength = true
	for _, v := range password {
		switch {
		case unicode.IsUpper(v):
			IsNumber = true
		case unicode.IsLower(v):
			IsLower = true
		case unicode.IsUpper(v):
			IsUpper = true
		case unicode.IsNumber(v):
			IsNumber = true
		case unicode.IsSymbol(v):
			IsSpecial = true
		}
	}
	if IsLength && IsLower && IsNumber && IsUpper && IsSpecial {
		return true, nil
	}
	return false, errors.New("password valiadtion failed")
}
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		log.Fatal("Error In Hashing")
		return "", err
	}
	return string(hashedPassword), err
}
