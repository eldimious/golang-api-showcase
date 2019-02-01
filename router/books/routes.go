package books

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/eldimious/golang-api-showcase/domain/books"
	domainErrors "github.com/eldimious/golang-api-showcase/domain/errors"
	"github.com/gin-gonic/gin"
)

// NewRoutesFactory create and returns a factory to create routes for the machines
func NewRoutesFactory(group *gin.RouterGroup) func(service books.BookService) {
	bookRoutesFactory := func(service books.BookService) {

		group.POST("/:authorId/books", func(c *gin.Context) {
			authorId, err := strconv.Atoi(c.Param("authorId"))
			fmt.Println(authorId)
			if err != nil {
				appError := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
				c.Error(appError)
				return
			}
			book, err := Bind(c, authorId)
			fmt.Println(*book)
			newBook, err := service.CreateBook(book)
			fmt.Println(*newBook)
			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusCreated, *toResponseModel(newBook))
		})

		group.GET("/:authorId/books/:bookId", func(c *gin.Context) {
			_, err := strconv.Atoi(c.Param("authorId"))
			if err != nil {
				appError := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
				c.Error(appError)
				return
			}

			id, err := strconv.Atoi(c.Param("bookId"))
			if err != nil {
				appError := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
				c.Error(appError)
				return
			}
			if err != nil {
				appError := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
				c.Error(appError)
				return
			}

			result, err := service.ReadBook(id)
			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, result)
		})
	}

	return bookRoutesFactory
}
