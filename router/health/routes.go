package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewRoutesFactory create and returns a factory to create routes to check API health
func NewRoutesFactory() func(group *gin.RouterGroup) {
	healthRoutesFactory := func(group *gin.RouterGroup) {
		group.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "health ok")
		})
	}

	return healthRoutesFactory
}
