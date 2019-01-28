package config

import (
	"github.com/eldimious/golang-api-showcase/utils/env"
)

// Config is a struct that contains configuration variables
type Config struct {
	Environment string
	Port        string
	Postgres    *Postgres
}

// Postgres is a struct that contains PostgreSQL configuration variables
type Postgres struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
}

// NewConfig creates a new Config struct
func NewConfig() (*Config, error) {
	env.CheckDotEnv()

	return &Config{
		Environment: env.MustGet("ENV"),
		Port:        env.MustGet("PORT"),
		Postgres: &Postgres{
			Host:     env.MustGet("POSTGRES_HOST"),
			Port:     env.MustGet("POSTGRES_PORT"),
			User:     env.MustGet("POSTGRES_USER"),
			DB:       env.MustGet("POSTGRES_DB"),
			Password: env.MustGet("POSTGRES_PASSWORD"),
		},
	}, nil
}
