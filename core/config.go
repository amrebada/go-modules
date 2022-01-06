package core

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	PORT     string
	HOST     string
	USER     string
	PASSWORD string
	DBNAME   string
}

type Config struct {
	InternalSecret string
	Port           int
	Database       DatabaseConfig
	JWT_SECRET     string
}

var config *Config

func NewConfig() *Config {
	godotenv.Load()
	if config == nil {
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			port = 8080
		}
		db_port := os.Getenv("DB_PORT")
		db_host := os.Getenv("DB_HOST")
		db_user := os.Getenv("DB_USER")
		db_password := os.Getenv("DB_PASSWORD")
		db_name := os.Getenv("DB_NAME")
		jwt_secret := os.Getenv("JWT_SECRET")

		return &Config{
			InternalSecret: "secret",
			Port:           port,
			Database: DatabaseConfig{
				PORT:     db_port,
				HOST:     db_host,
				USER:     db_user,
				PASSWORD: db_password,
				DBNAME:   db_name,
			},
			JWT_SECRET: jwt_secret,
		}
	}
	return config
}
