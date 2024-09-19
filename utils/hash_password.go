package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// generate a hash from brcypt
	bytesHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytesHash), nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	// check if it is a valid password
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))
	return err == nil
}
