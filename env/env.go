package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Setup loads the .env file
func Setup() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Get returns the value of the environment variable key

func Get(key string) string {
	return os.Getenv(key)
}
