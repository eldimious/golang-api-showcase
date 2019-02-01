package main

import (
	"net/http"

	"github.com/eldimious/golang-api-showcase/config"
	postgres "github.com/eldimious/golang-api-showcase/data/database"

	"github.com/eldimious/golang-api-showcase/domain/authors"
	"github.com/eldimious/golang-api-showcase/domain/books"
	"github.com/eldimious/golang-api-showcase/router"

	authorsStore "github.com/eldimious/golang-api-showcase/data/authors"
	booksStore "github.com/eldimious/golang-api-showcase/data/books"
)

func main() {
	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	postgres, err := postgres.Connect(configuration.Postgres)
	if err != nil {
		panic(err)
	}
	booksRepo := booksStore.New(postgres)
	booksSvc := books.NewService(booksRepo)

	authorsRepo := authorsStore.New(postgres, booksRepo)
	booksRepo.AddAuthorFKConstraint() // Author relation must exist before we add the constraint
	authorsSvc := authors.NewService(authorsRepo)
	httpRouter := router.NewHTTPHandler(authorsSvc, booksSvc)
	err = http.ListenAndServe(":"+configuration.Port, httpRouter)
	if err != nil {
		panic(err)
	}

	defer postgres.Close()
}
