package main

import (
	"github.com/eldimious/golang-api-showcase/config"
	postgres "github.com/eldimious/golang-api-showcase/data/database"
)

func main() {
	configuration := config.NewConfig()
	if err != nil {
		panic(err)
	}
	postgres, err := postgres.Open(configuration.Postgres)
	if err != nil {
		panic(err)
	}
}
