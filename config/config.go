package config

import (
	"log"
	"os"

	"github.com/joho/godotenv" //go ka package hai which helps load env variables from .env file file into the application.
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
