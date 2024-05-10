package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	DatabaseUrl   string
}

// Function to load config
func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		serverAddress = "8080"
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return nil, errors.New("Missing DatabaseUrl")
	}
	return &Config{
		ServerAddress: serverAddress,
		DatabaseUrl:   databaseUrl,
	}, nil

}
