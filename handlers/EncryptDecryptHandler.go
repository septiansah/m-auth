package handlers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CompareHashedAndSalted(hashedVal string, plainVal []byte) (bool, error) {
	byteHash := []byte(hashedVal)
	
	if err := bcrypt.CompareHashAndPassword(byteHash, plainVal); err != nil {
		return false, fmt.Errorf("Error compare hash and salt process, reason: " + err.Error())
	}

	return true, nil
}