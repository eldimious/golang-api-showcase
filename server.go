package main

import (
	"net/http"

	"github.com/eldimious/golang-api-showcase/config"
	db "github.com/eldimious/golang-api-showcase/data/database"

	"github.com/eldimious/golang-api-showcase/domain/authors"
	"github.com/eldimious/golang-api-showcase/domain/books"
	router "github.com/eldimious/golang-api-showcase/router/http"

	authorsStore "github.com/eldimious/golang-api-showcase/data/authors"
	booksStore "github.com/eldimious/golang-api-showcase/data/books"
)

func main() {
	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	// establish DB connection
	db, err := db.Connect(configuration.Postgres)
	if err != nil {
		panic(err)
	}
	booksRepo := booksStore.New(db)
	booksSvc := books.NewService(booksRepo)

	authorsRepo := authorsStore.New(db, booksRepo)
	booksRepo.AddAuthorFKConstraint() // Author relation must exist before we add the constraint
	authorsSvc := authors.NewService(authorsRepo)
	httpRouter := router.NewHTTPHandler(authorsSvc, booksSvc)
	err = http.ListenAndServe(":"+configuration.Port, httpRouter)
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
