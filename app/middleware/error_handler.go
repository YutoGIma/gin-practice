package middleware

import (
	"myapp/app/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			switch e := err.(type) {
			case *errors.AppError:
				c.JSON(e.Code, gin.H{
					"error": e.Message,
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}
	}
}
