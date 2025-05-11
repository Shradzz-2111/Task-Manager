package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func HashPassword(password string) (string,error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil 
}

func VerifyPassword(userPassword string,providedPassword string) error{
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
}