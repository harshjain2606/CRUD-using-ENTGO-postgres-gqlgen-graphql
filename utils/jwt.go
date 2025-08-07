package utils


import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_secret_key") // ideally from env

func GenerateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
