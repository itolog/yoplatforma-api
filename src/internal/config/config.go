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

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		logging.Fatal("Error loading .env file")
	}

	return &Config{
		Port: os.Getenv("APP_PORT"),
		Host: os.Getenv("APP_HOST"),
	}
}
