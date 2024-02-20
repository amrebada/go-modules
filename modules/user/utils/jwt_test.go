package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/amrebada/go-modules/modules/user/config"
	"github.com/stretchr/testify/assert"
)

func TestTokenUtils(t *testing.T) {
	data := map[string]any{
		"username": "john",
		"role":     "admin",
	}

	userConfig := &config.UserModuleConfiguration{
		Authentication: config.AuthenticationConfiguration{
			JwtSecret:                      "secret",
			JwtExpirationDurationInSeconds: 60 * 60,
		},
	}

	t.Run("it should generate token", func(t *testing.T) {

		token, err := GenerateToken(data, userConfig)
		fmt.Println(token, err)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("it should parse token ", func(t *testing.T) {

		token, err := GenerateToken(data, userConfig)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		claims, err := ParseToken(token, userConfig)
		assert.NoError(t, err)
		assert.NotEmpty(t, claims)
		assert.Equal(t, "john", claims["username"])
		assert.Equal(t, "admin", claims["role"])
	})

	t.Run("it should validate and return success", func(t *testing.T) {

		token, err := GenerateToken(data, userConfig)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		isValid, err := ValidateToken(token, userConfig)
		assert.NoError(t, err)
		assert.True(t, isValid)
	})

	t.Run("it should validate & return error once expired", func(t *testing.T) {
		userConfig.Authentication.JwtExpirationDurationInSeconds = 1
		token, err := GenerateToken(data, userConfig)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		time.Sleep(time.Second * 2)
		isValid, err := ValidateToken(token, userConfig)
		assert.Error(t, err)
		assert.False(t, isValid)
	})

}
