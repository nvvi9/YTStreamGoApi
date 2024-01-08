package utils

import (
	"YTStreamGoApi/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateJWT(id uuid.UUID, config *config.Config) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(config.JwtExpiresIn).Unix()

	return token.SignedString([]byte(config.JwtSecret))
}
