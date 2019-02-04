package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	local      = "local"
	staging    = "staging"
	production = "production"
	env        = "ENV"
)

// MustGet will return the env or panic if not present.
func MustGet(key string) string {
	val := os.Getenv(key)
	if val == "" && key != "PORT" {
		panic("Env key missing " + key)
	}
	return val
}

// CheckDotEnv loads environment variables from .env file for development environment
func CheckDotEnv() {
	err := godotenv.Load()
	if err != nil && os.Getenv(env) == local {
		log.Println("Error loading .env file")
	}
}
