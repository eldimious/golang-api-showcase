package router

import (
	"net/http"

	"github.com/eldimious/golang-api-showcase/domain/authors"
	"github.com/eldimious/golang-api-showcase/domain/books"

	authorsRoutes "github.com/eldimious/golang-api-showcase/router/authors"
	errors "github.com/eldimious/golang-api-showcase/router/errors"
	healthRoutes "github.com/eldimious/golang-api-showcase/router/health"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewHTTPHandler returns the HTTP requests handler
func NewHTTPHandler(authorSvc authors.AuthorService, booksSvc books.BookService) http.Handler {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	router.Use(errors.Handler)

	healthGroup := router.Group("/health")
	healthRoutes.NewRoutesFactory()(healthGroup)

	api := router.Group("/api")
	//api.Use(authMiddleware)

	authorsGroup := api.Group("/authors")
	authorsRoutes.NewRoutesFactory(authorsGroup)(authorSvc)

	return router
}
