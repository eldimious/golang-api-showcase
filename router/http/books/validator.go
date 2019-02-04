package books

import (
	"github.com/eldimious/golang-api-showcase/domain/books"
	"github.com/gin-gonic/gin"
)

// BookValidator is a struct used to validate the JSON payload representing a book.
type BookValidator struct {
	Name      string `binding:"required" json:"name"`
	Publisher string `binding:"required" json:"publisher"`
}

// Bind validates the JSON payload and returns a Book instance corresponding to the payload.
func Bind(c *gin.Context, authorId int) (*books.Book, error) {
	var json BookValidator
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}

	book := &books.Book{
		Name:      json.Name,
		Publisher: json.Publisher,
		AuthorId:  authorId,
	}

	return book, nil
}
