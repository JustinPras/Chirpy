package auth

import (
	"testing"
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