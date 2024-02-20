package utils

import (
	"time"

	"github.com/amrebada/go-modules/modules/user/config"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(data map[string]any, config *config.UserModuleConfiguration) (string, error) {
	secret := config.Authentication.JwtSecret
	token := jwt.New(jwt.GetSigningMethod("HS512"))
	claims := token.Claims.(jwt.MapClaims)
	for key, value := range data {
		claims[key] = value
	}
	claims["exp"] = time.Now().Add(time.Duration(config.Authentication.JwtExpirationDurationInSeconds * int(time.Second))).Unix()
	return token.SignedString([]byte(secret))
}

func ParseToken(tokenString string, config *config.UserModuleConfiguration) (jwt.MapClaims, error) {
	secret := config.Authentication.JwtSecret
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}

func ValidateToken(tokenString string, config *config.UserModuleConfiguration) (bool, error) {
	_, err := ParseToken(tokenString, config)
	if err != nil {
		return false, err
	}
	return true, nil
}
