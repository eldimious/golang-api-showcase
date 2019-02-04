package config

import (
	"github.com/eldimious/golang-api-showcase/utils/env"
)

// Config is a struct that contains configuration variables
type Config struct {
	Environment string
	Port        string
	Database    *Database
}

// Database is a struct that contains DB's configuration variables
type Database struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
}

// NewConfig creates a new Config struct
func NewConfig() (*Config, error) {
	env.CheckDotEnv()
	port := env.MustGet("PORT")

	if port == "" {
		port = "3000"
	}
	return &Config{
		Environment: env.MustGet("ENV"),
		Port:        port,
		Database: &Database{
			Host:     env.MustGet("DATABASE_HOST"),
			Port:     env.MustGet("DATABASE_PORT"),
			User:     env.MustGet("DATABASE_USER"),
			DB:       env.MustGet("DATABASE_DB"),
			Password: env.MustGet("DATABASE_PASSWORD"),
		},
	}, nil
}
