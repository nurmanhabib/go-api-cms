package middleware

import (
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"go-api-cms/infrastructure/exception"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Log panic and stack trace
				log.Printf("[PANIC] %v\nStack Trace:\n%s\n", r, string(debug.Stack()))
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
				c.Abort()
			}
		}()

		c.Next()

		errLast := c.Errors.Last()
		if errLast != nil {
			var svcException *exception.ServiceError
			if errors.As(errLast, &svcException) {
				svcException.Gin(c)
				return
			}
		}

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				// Log error with stack trace
				log.Printf("[ERROR] %v\nStack Trace:\n%s\n", err.Err, string(debug.Stack()))
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			c.Abort()
		}
	}
}
