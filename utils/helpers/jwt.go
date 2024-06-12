package utils

import (
	"os"
	"time"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

var JwtExpTime int = 3

func GenerateToken(email string, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email": email,
		"exp": time.Now().Add(time.Hour * 3).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func VerifyToken(token string) (string, string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return "", "", errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return "", "", errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return "", "", errors.New("invalid token claims")
	}

	userId := string(claims["userId"].(string))
	email := string(claims["email"].(string))

	return userId, email, nil
}