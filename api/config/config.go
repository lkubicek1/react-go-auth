package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvVars struct {
	AUTH0_DOMAIN   string `mapstructure:"AUTH0_DOMAIN"`
	AUTH0_AUDIENCE string `mapstructure:"AUTH0_AUDIENCE"`
	PORT           string `mapstructure:"PORT"`
}

func LoadEnv() EnvVars {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Map environment variables to the struct
	return EnvVars{
		AUTH0_DOMAIN:   os.Getenv("AUTH0_DOMAIN"),
		AUTH0_AUDIENCE: os.Getenv("AUTH0_AUDIENCE"),
		PORT:           os.Getenv("PORT"),
	}
}
