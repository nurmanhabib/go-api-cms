package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceError struct {
	StatusCode int
	Message    string
	Violations []FieldViolation
}

func NewServiceError(statusCode int, message string) *ServiceError {
	return &ServiceError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func NewUnprocessableEntity(violations []FieldViolation) *ServiceError {
	return &ServiceError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Unprocessable Entity",
		Violations: violations,
	}
}

func (e *ServiceError) Error() string {
	return e.Message
}

func (e *ServiceError) Gin(c *gin.Context) {
	if len(e.Violations) > 0 {
		c.JSON(e.StatusCode, map[string]interface{}{
			"code":       e.StatusCode,
			"message":    e.Message,
			"violations": e.Violations,
		})
		return
	}

	c.JSON(e.StatusCode, e.Message)
}
