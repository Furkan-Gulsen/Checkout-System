package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceError struct {
	Message string
	Code    int
}

func (e *ServiceError) Error() string {
	return e.Message
}

func (e *ServiceError) GetCode() int {
	return e.Code
}

func NewServiceError(message string, code int, data *interface{}) error {
	return &ServiceError{
		Message: message,
		Code:    code,
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if err := c.Errors.Last(); err != nil {
			if serviceErr, ok := err.Err.(*ServiceError); ok {
				errorCode := serviceErr.GetCode()
				c.JSON(errorCode, gin.H{"message": serviceErr.Message})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		}
	}
}
