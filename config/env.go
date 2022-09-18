package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func cockroachURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("connectionString")
}
