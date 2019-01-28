package router

import (
	"net/http"

	domain "github.com/eldimious/golang-api-showcase/domain/errors"
	"github.com/gin-gonic/gin"
)

// Handler is Gin middleware to handle errors.
func Handler(c *gin.Context) {
	// Execute request handlers and then handle any errors
	c.Next()

	errs := c.Errors

	if len(errs) > 0 {
		err, ok := errs[0].Err.(*domain.AppError)
		if ok {
			switch err.Type {
			case domain.NotFound:
				c.JSON(http.StatusNotFound, err.Error())
				return
			case domain.ValidationError:
				c.JSON(http.StatusBadRequest, err.Error())
				return
			case domain.ResourceAlreadyExists:
				c.JSON(http.StatusConflict, err.Error())
				return
			case domain.NotAuthenticated:
				c.JSON(http.StatusUnauthorized, err.Error())
				return
			case domain.NotAuthorized:
				c.JSON(http.StatusForbidden, err.Error())
				return
			case domain.RepositoryError:
			default:
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		// Error is not AppError, return a generic internal server error
		c.JSON(http.StatusInternalServerError, "Internal Server Errror")
		return
	}
}
