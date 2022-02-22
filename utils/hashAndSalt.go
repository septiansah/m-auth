package utils

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(input []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(input, bcrypt.MinCost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}