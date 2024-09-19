package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = os.Getenv("TOKEN_SECRET_KEY")

func GenerateToken(email string, userID string) (string, error) {
	// generate a new token with information about email, userID and expiration
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func ValidateToken(token string) (string, error) {
	// decode the token passed and validate if its a correct token or not
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("wrong signing Method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", errors.New("could not parse token")
	}
	if !parsedToken.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("wrong data in token")
	}
	//email := claims["email"].(string)
	userID := claims["userID"].(string)
	return userID, nil
}
