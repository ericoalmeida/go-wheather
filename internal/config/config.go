package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GeoapifyBaseUrl string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	geoapifyBaseUrl := os.Getenv("GEOAPIFY_BASE_URL")

	if geoapifyBaseUrl == "" {
		log.Fatal("Environment variabels not defined.")
	}

	return &Config{
		GeoapifyBaseUrl: geoapifyBaseUrl,
	}
}
