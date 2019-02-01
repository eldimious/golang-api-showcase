package authors

import (
	"github.com/eldimious/golang-api-showcase/domain/authors"
	"github.com/gin-gonic/gin"
)

// AuthorValidator is a struct used to validate the JSON payload representing an author.
type AuthorValidator struct {
	Name    string `binding:"required" json:"name"`
	Surname string `binding:"required" json:"surname"`
}

// Bind validates the JSON payload and returns a Author instance corresponding to the payload.
func Bind(c *gin.Context) (*authors.Author, error) {
	var json AuthorValidator
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}

	author := &authors.Author{
		Name:    json.Name,
		Surname: json.Surname,
	}

	return author, nil
}
