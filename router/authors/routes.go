package authors

import (
	"net/http"
	"strconv"

	"github.com/eldimious/golang-api-showcase/domain/authors"
	domainErrors "github.com/eldimious/golang-api-showcase/domain/errors"
	"github.com/gin-gonic/gin"
)

// NewRoutesFactory create and returns a factory to create routes for the authors
func NewRoutesFactory(group *gin.RouterGroup) func(service authors.AuthorService) {
	authorRoutesFactory := func(service authors.AuthorService) {
		group.POST("/", func(c *gin.Context) {
			author, err := Bind(c)
			if err != nil {
				appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
				c.Error(appError)
				return
			}

			created, err := service.CreateAuthor(author)
			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusCreated, created)
		})

		group.GET("/:authorId", func(c *gin.Context) {
			id := c.Param("authorId")
			i, err := strconv.Atoi(id)
			result, err := service.ReadAuthor(i)
			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, result)
		})
	}

	return authorRoutesFactory
}
