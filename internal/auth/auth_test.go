package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("Password123!")
	if hash == "" || err != nil {
		t.Errorf("Error Hashing Password: %v", err)
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "Password123!"
	hash, _ := HashPassword(password)
	err := CheckPasswordHash(hash, password)
	if err != nil {
		t.Errorf("Hash does not match password: %v", err)
	}
}

func TestMakeJWT(t *testing.T) {
	userID, _ := uuid.Parse("6b61fa75-e753-4e95-bc5b-c8aa9c629e1a")
	const expiresIn = 5 * time.Minute
	const tokenSecret = "7Lqz!pD@f9$e3Gx^KbJ6Wu#cZrMqT1vN"
	
	_, err := MakeJWT(userID, tokenSecret, expiresIn)
	if err != nil {
		t.Errorf("Error making JWT: %v", err)
	}
}

func TestValidateJWT(t *testing.T) {
	userID, _ := uuid.Parse("6b61fa75-e753-4e95-bc5b-c8aa9c629e1a")
	const expiresIn = 5 * time.Minute
	const tokenSecret = "7Lqz!pD@f9$e3Gx^KbJ6Wu#cZrMqT1vN"
	
	tokenString, err := MakeJWT(userID, tokenSecret, expiresIn)
	if err != nil {
		t.Errorf("Error making JWT: %v", err)
	}

	userIDActual, err := ValidateJWT(tokenString, tokenSecret)
	if err != nil {
		t.Errorf("Error validating JWT: %v", err)
	}

	if userID != userIDActual {
		t.Errorf("Expecting: %v\nActual: %v", userID, userIDActual)
	}
}