package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Param handling

func GetParam(c *gin.Context, param string) (string, error) {
	value, exists := c.Params.Get(param)
	if !exists {
		return "", CustomError{http.StatusBadRequest, "Missing parameter: " + param}
	}

	return value, nil
}

// Error handling

type CustomError struct {
	StatusCode int
	Message    string
}

func (ce CustomError) Error() string {
	return ce.Message
}

func Fatal(code int, err error) {
	panic(CustomError{code, err.Error()})
}

func CheckFatal(checked error, code int, err error) {
	if checked != nil {
		panic(CustomError{code, err.Error()})
	}
}

func RecoveryHandlers() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if customErr, ok := err.(CustomError); ok {
					// Log the custom error message
					log.Println("Custom Error:", customErr)

					// Return the custom error as the response
					c.JSON(customErr.StatusCode, gin.H{
						"error": customErr.Message,
					})
				} else {
					// Log the panic message
					log.Println("Panic:", err)

					// For other panics, return a generic error response
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": "Internal Server Error",
					})
				}

				// Abort the current request
				c.Abort()
			}
		}()

		// Continue processing the request
		c.Next()
	}
}
