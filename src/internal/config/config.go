package config

import (
	"api_platforma/src/pkg/logging"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port string `env:"PORT"`
	Host string `env:"HOST"`
}

const (
	DEV  = "development"
	PROD = "production"
)

func NewConfig() *Config {
	appEnv := os.Getenv("FOO_ENV")

	if appEnv == PROD {
		err := godotenv.Load(".env")
		if err != nil {
			logging.Fatal("Error loading .env file")
		}
	} else if appEnv == DEV {
		err := godotenv.Load(".env.development")
		if err != nil {
			logging.Fatal("Error loading .env.development file")
		}
	} else {
		err := godotenv.Load(".env.development.local")
		if err != nil {
			logging.Fatal("Error loading .env.development.local file")
		}
	}

	return &Config{
		Port: os.Getenv("APP_YP_PORT"),
		Host: os.Getenv("APP_YP_HOST"),
	}
}
