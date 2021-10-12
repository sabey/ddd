package ddd

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var hmacSecret []byte = []byte("abcdefghijklmnopqrstuvwxyz")

func SignJWTClaims(
	email string,
) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"nbf":   time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, _ := token.SignedString(hmacSecret)

	return tokenString
}

func ParseJWTClaims(
	tokenString string,
) string {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSecret, nil
	})
	if err != nil {
		// invalid token
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if _, ok := claims["email"]; !ok {
			// email doesn't exist
			return ""
		}

		return claims["email"].(string)
	}

	return ""
}
