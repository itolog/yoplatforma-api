package config

import (
	"api_platforma/src/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port string `env:"PORT" env-default:":8000"`
	Host string `env:"HOST" env-default:"localhost"`
}

const (
	DEV  = "development"
	PROD = "production"
)

func NewConfig() *Config {
	appEnv := os.Getenv("APP_YP_ENV")

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
	}

	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		logging.Warn(err)
	}

	return &Config{
		Port: cfg.Port,
		Host: cfg.Host,
	}
}
