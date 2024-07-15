package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint) (string, error) {
	var claims = jwt.MapClaims{}
	claims["id"] = userID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	var process = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := process.SignedString([]byte("passkeyJWT"))
	if err != nil {
		return "", err
	}
	return result, nil
}

func DecodeToken(token *jwt.Token) uint {
	var claims = token.Claims.(jwt.MapClaims)
	var userID = claims["id"].(float64)
	return uint(userID)
}