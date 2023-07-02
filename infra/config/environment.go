package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	ambiente := os.Getenv("GO_ENV")

	if ambiente == "" {
		loadEnvFile()
	}

}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
