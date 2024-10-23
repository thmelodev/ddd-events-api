package httpServer

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			errorType := reflect.TypeOf(err).Elem().Name()

			switch e := err.(type) {
			case *apiErrors.RepositoryError, *apiErrors.InvalidPropsError:
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   errorType,
					"message": e.Error(),
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "InternalServerError",
					"message": err.Error(),
				})
			}

			c.Abort()
		}
	}
}
