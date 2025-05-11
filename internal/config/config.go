package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GeoapifyBaseUrl string
}

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Aviso: arquivo .env nao encontrado. Assumindo valores do ambiente.")
	}
}

func GetEnv(key string) string {
	envValue := os.Getenv(key)

	if envValue == "" {
		log.Fatal("Environment variabels not defined.")
	}

	return envValue
}
