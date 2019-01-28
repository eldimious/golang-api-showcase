package main

import (
	"fmt"

	"github.com/eldimious/golang-api-showcase/config"
	postgres "github.com/eldimious/golang-api-showcase/data/database"
)

func main() {
	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	postgres, err := postgres.Open(configuration.Postgres)
	if err != nil {
		panic(err)
	}
	fmt.Println(postgres)
}
